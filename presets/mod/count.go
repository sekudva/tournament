package mod

import (
	"github.com/sekudva/strategika/internal/domain"
)

// Doubler
func Doubler() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			ctx.ModState[domain.RepeatCounter] = 0
			return core
		}

		if ctx.ModState[domain.RepeatCounter] == 0 {
			ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		}

		ctx.ModState[domain.RepeatCounter]++

		if ctx.ModState[domain.RepeatCounter] <= 2 {
			return domain.Act(ctx.ModState[domain.ActRecorder])
		}

		ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		ctx.ModState[domain.RepeatCounter] = 1
		return domain.Act(ctx.ModState[domain.ActRecorder])
	}
}

// Дважды повторяет только Share и Take
func JournalistMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			ctx.ModState[domain.RepeatCounter] = 0
			return core
		}

		if ctx.ModState[domain.RepeatCounter] == 0 {
			ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())
		}

		ctx.ModState[domain.RepeatCounter]++

		if ctx.ModState[domain.RepeatCounter] <= 2 {
			if domain.Act(ctx.ModState[domain.ActRecorder]) != domain.Hold {
				return domain.Act(ctx.ModState[domain.ActRecorder])
			}
		}

		ctx.ModState[domain.ActRecorder] = int(ctx.History.OpLastAct())

		if domain.Act(ctx.ModState[domain.ActRecorder]) != domain.Hold {
			ctx.ModState[domain.RepeatCounter] = 1
			return domain.Act(ctx.ModState[domain.ActRecorder])
		}

		return core
	}
}

func GoByMajorityMod() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {

		var shareCount, holdCount, takeCount int
		for _, round := range ctx.History {
			switch round.OpAct {
			case domain.Share:
				shareCount++
			case domain.Hold:
				holdCount++
			case domain.Take:
				takeCount++
			}
		}

		// Большинство Share → Share
		if shareCount > holdCount && shareCount > takeCount {
			return domain.Share
		}

		// Большинство Take → Take
		if takeCount > shareCount && takeCount > holdCount {
			return domain.Take
		}

		// Большинство Hold → Hold
		if holdCount >= shareCount && holdCount >= takeCount {
			return domain.Hold
		}

		return core
	}
}

func Exploiter() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		n := len(ctx.History)

		if _, ok := ctx.ModState[domain.PunishCounter]; !ok {
			ctx.ModState[domain.PunishCounter] = 0
		}
		if _, ok := ctx.ModState[domain.SafeCounter]; !ok {
			ctx.ModState[domain.SafeCounter] = 0
		}

		if n >= 1 {
			myPrev := ctx.History.My2LastAct()
			op := ctx.History.OpLastAct()

			if myPrev != domain.Share {
				if op == domain.Take {
					ctx.ModState[domain.PunishCounter]++
				} else {
					ctx.ModState[domain.SafeCounter]++
				}
			}

		}

		if n <= 10 {
			return core
		}

		myPrev := ctx.History.My2LastAct()
		my := ctx.History.MyLastAct()
		op := ctx.History.OpLastAct()

		punish, safe := ctx.ModState[domain.PunishCounter], ctx.ModState[domain.SafeCounter]

		switch myPrev {
		case domain.Take:
			if op != domain.Take || safe > punish {
				return domain.Take
			}
			return domain.Hold

		case domain.Hold:
			if op != domain.Take {
				return domain.Share
			}
			return domain.Hold

		case domain.Share:
			if op != domain.Take && my != domain.Take {
				return domain.Take
			} else if punish > (safe*4) || my == domain.Take {
				return domain.Hold
			}

			return domain.Share

		default:
			return core
		}

	}
}

// BimodalPredictor — стратегия, использующая 2-битный насыщающий счётчик
// для предсказания хода оппонента.
func BPU() domain.Modifier {
	// Счётчики для каждого возможного действия оппонента
	// 0–1: предсказываем, что оппонент НЕ сделает это действие
	// 2–3: предсказываем, что оппонент СДЕЛАЕТ это действие
	counters := map[domain.Act]int{
		domain.Share: 2, // изначально предсказываем Share
		domain.Hold:  1, // изначально не предсказываем Hold
		domain.Take:  1, // изначально не предсказываем Take
	}

	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) == 0 {
			return domain.Share
		}

		// 1. Предсказываем следующий ход оппонента
		predicted := predictNext(counters)

		// 2. Выбираем лучший ответ на предсказание
		bestResponse := getBestResponse(predicted)

		// 3. Обновляем счётчики на основе реального хода оппонента
		actual := ctx.History.OpLastAct()
		updateCounters(counters, actual)

		return bestResponse
	}
}

// predictNext выбирает действие с наибольшим счётчиком
func predictNext(counters map[domain.Act]int) domain.Act {
	var bestAct domain.Act
	maxCount := -1
	for act, count := range counters {
		if count > maxCount {
			maxCount = count
			bestAct = act
		}
	}
	return bestAct
}

// updateCounters обновляет счётчики по принципу насыщения
func updateCounters(counters map[domain.Act]int, actual domain.Act) {
	for act := range counters {
		if act == actual {
			// Предсказание сбылось → увеличиваем (насыщение до 3)
			if counters[act] < 3 {
				counters[act]++
			}
		} else {
			// Предсказание не сбылось → уменьшаем (насыщение до 0)
			if counters[act] > 0 {
				counters[act]--
			}
		}
	}
}

// getBestResponse — оптимальный ответ на ожидаемый ход оппонента.
func getBestResponse(predicted domain.Act) domain.Act {
	switch predicted {
	case domain.Share:
		return domain.Take // если оппонент Share, лучшее — Take
	case domain.Hold:
		return domain.Share // против Hold — Share
	case domain.Take:
		return domain.Share // против Take — Hold? Но в классике лучше Share
	default:
		return domain.Share
	}
}
