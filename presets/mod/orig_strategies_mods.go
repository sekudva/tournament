package mod

import (
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// Модификатор для стратегии Туллока, накопительная стратегия
// Снижение кооперации по накоплению предательств
func TullockMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.Round <= 11 {
			return core
		}

		// Смотрим последние 10 ходов оппонента
		start := max(0, len(ctx.History)-10)
		coopCount := 0
		for i := start; i < len(ctx.History); i++ {
			if ctx.History[i].OpAct == domain.Share {
				coopCount++
			}
		}

		coopRate := float64(coopCount) / 10.0

		// Если 9 из 10 коопераций — щедрый шаг
		if coopCount >= 9 {
			if rand.Float64() < 0.8 {
				return domain.Share
			} else {
				return domain.Take
			}
		}

		// Иначе: кооперируем на (coopRate - 0.1)%
		prob := max(0, coopRate-0.1)
		if rand.Float64() < prob {
			return domain.Share
		} else {
			return domain.Take
		}

	}
}

func AdamsMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		lastOp := ctx.History.OpLastAct()

		// Инициализация
		if _, ok := ctx.ModState[domain.CurrThresholdCounter]; !ok {
			ctx.ModState[domain.CurrThresholdCounter] = 4 // am — текущий порог
			ctx.ModState[domain.RepeatCounter] = 0        // mc — счётчик дефектов
		}

		am := ctx.ModState[domain.CurrThresholdCounter]
		mc := ctx.ModState[domain.RepeatCounter]

		// Считаем дефекты
		if lastOp == domain.Take {
			mc++
			ctx.ModState[domain.RepeatCounter] = mc
		}

		// Если дефектов меньше порога → Share
		if float64(mc) < float64(am) {
			return domain.Share
		}

		// Если ровно достигли порога → один раз Take
		if float64(mc) == float64(am) {
			ctx.ModState[domain.RepeatCounter] = 0
			return domain.Take
		}

		// Превысили порог → уменьшаем порог, даём шанс
		am = am / 2
		ctx.ModState[domain.CurrThresholdCounter] = am
		ctx.ModState[domain.RepeatCounter] = 0

		if rand.Float64() < float64(am)/4.0 {
			return domain.Share
		}
		return domain.Take
	}
}

func EatherleyMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		lastOp := ctx.History.OpLastAct()

		// Инициализация
		if _, ok := ctx.ModState[domain.GoodStreakCounter]; !ok {
			ctx.ModState[domain.GoodStreakCounter] = 0 // nj — кооперации оппонента
		}

		// Если оппонент только что дефектил → Take
		if lastOp == domain.Take {
			return core
		}

		// Считаем кооперации оппонента
		if lastOp == domain.Share {
			ctx.ModState[domain.GoodStreakCounter]++
		}

		// Вероятность кооперации = доля коопераций оппонента
		n := len(ctx.History)
		p := float64(ctx.ModState[domain.GoodStreakCounter]) / float64(n)
		if rand.Float64() < p {
			return core
		}

		return domain.Take
	}
}

func CaveMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		n := len(ctx.History)
		if n == 0 {
			return domain.Share
		}

		// Инициализация счётчика дефектов
		if _, ok := ctx.ModState[domain.DefectSum]; !ok {
			ctx.ModState[domain.DefectSum] = 0
		}

		lastOp := ctx.History.OpLastAct()
		if lastOp == domain.Take {
			ctx.ModState[domain.DefectSum]++
		}
		defectSum := ctx.ModState[domain.DefectSum]
		defectPc := defectSum * 100 / n

		// Проверки на безумного оппонента
		if (n > 19 && defectPc > 79) ||
			(n > 29 && defectPc > 65) ||
			(n > 39 && defectPc > 39) {
			return domain.Take
		}

		// Если дефектов мало — случайный ответ
		// (в агенте должна быть стратегия рандома)
		if defectSum <= 17 {
			return core
		}

		// Иначе всегда Take
		return domain.Take
	}
}

func ChampionMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		n := len(ctx.History)
		if n == 0 {
			ctx.ModState[domain.CoopSum] = 0
			return domain.Share
		}

		if ctx.History.OpLastAct() == domain.Share {
			ctx.ModState[domain.CoopSum]++
		}

		// Ходы 11-25 — TFT (у агента должно быть TFT!)
		if n <= 25 {
			if n <= 10 {
				return domain.Share
			}
			return core
		}

		// После 25 — условный TFT(с прощением)
		coopRat := float64(ctx.ModState[domain.CoopSum]) / float64(n)

		if ctx.History.OpLastAct() == domain.Take &&
			coopRat < 0.6 && coopRat <= rand.Float64() {
			return domain.Take
		}

		return domain.Share
	}
}

// на 2 Take прощает 75%, на 1 предыдущий Take но следующий не Take - 100%
func LeyvrazMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		op := ctx.History.OpLastAct()      // последний
		opPrev := ctx.History.Op2LastAct() // предпоследний

		if opPrev == domain.Take && !(op == domain.Take && rand.Float64() < 0.25) {
			return domain.Take
		}

		return domain.Share
	}
}

// PavlovMod — Win-Stay, Lose-Shift для трёх действий.
// Для базовой минусовой матрицы
func Pavlov() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		my := ctx.History.MyLastAct()
		op := ctx.History.OpLastAct()
		myPayoff, _ := domain.Payoff(my, op)

		if myPayoff >= 0 {
			return my
		}

		switch my {
		case domain.Share:
			return domain.Take
		default:
			// в остальных случаях смена на Share
			return domain.Share
		}
	}
}
