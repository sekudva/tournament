package mod

import (
	"math/rand/v2"

	"github.com/sekudva/strategika/internal/domain"
)

// WRAP MODS - вспомогательные надстройки над модификаторами

// Обертка для любого модификатора который смотрит на предыдущее действие оппонента
// Игнорирование применения модификатора если оппонент сделал Hold
func IgnoreHold(next domain.Modifier) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) > 0 && (ctx.History.OpLastAct() == domain.Hold) {
			return core
		}
		return next(core, ctx)
	}
}

// Применяет модификатор с веротяностью
func WithProbability(prob float64, next domain.Modifier) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if rand.Float64() < prob {
			return next(core, ctx)
		}
		return core
	}
}

func Sleep(awakeRound int, next domain.Modifier) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if len(ctx.History) < awakeRound {
			return core
		}
		return next(core, ctx)
	}
}

// возвращает мод после определенного количества действий во всей истори
func AfterTotal(act domain.Act, n int, next domain.Modifier, cnt domain.Counter) domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if _, ok := ctx.ModState[cnt]; !ok {
			ctx.ModState[cnt] = 0
		}

		if len(ctx.History) > 0 && ctx.History.OpLastAct() == act {
			ctx.ModState[cnt]++
		}

		if ctx.ModState[cnt] >= n {
			return next(core, ctx)
		}

		return core
	}
}
