package mod

import "github.com/sekudva/strategika/internal/domain"

// STRATEGY MODIFIERS FOR ARENA AND TRIAL LIST (1/Group and Arena)
// and REPUTATION-BASED STRATEGY MODIFIERS

// SeekWeak - ищет слабых и эксплуатирует
func SeekWeak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.6 && ctx.OpRep.Def < 0.35 {
			return domain.Take
		}
		return core
	}
}

// SeekFreak - ищет недружелюбных защищающихся и не взаимодействует
func SeekFreak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.3 && ctx.OpRep.Def > 0.65 {
			return domain.Hold
		}
		return core
	}
}

// SeekSneak - ищет пассивных
func SeekSneak() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.2 && ctx.OpRep.Def < 0.2 {
			return domain.Take
		}
		return core
	}
}

// SeekBrick - находит защищающихся дружелюбных и кооперирует
func SeekBrick() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.6 && ctx.OpRep.Def > 0.7 {
			return domain.Take
		}
		return core
	}
}

// Воин - ищет агрессивных и воюет с ними
func Warrior() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop < 0.15 && ctx.OpRep.Def > 0.75 {
			return domain.Take
		}
		return core
	}
}

// Кооперирует со слабыми
func Assistant() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.6 && ctx.OpRep.Def < 0.3 {
			return domain.Share
		}
		return core
	}
}

// Кооперирует даже если случайно вошел в цикл зла
func Mature() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Coop > 0.7 {
			return domain.Share
		}
		return core
	}
}

// Драчун, дерется с сильными
func Brawler() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Def > 0.9 {
			return domain.Take
		}
		return core
	}
}

// ищет слабых и тянет их на дно
func Quicksand() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		if ctx.OpRep.Def < 0.35 {
			return domain.Take
		}
		return core
	}
}

// PavlovMod — Win-Stay, Lose-Shift для трёх действий.
func EvilPavlov() domain.Modifier {
	return func(core domain.Act, ctx domain.ModContext) domain.Act {
		my := ctx.History.MyLastAct()
		op := ctx.History.OpLastAct()
		myPayoff, opPayoff := domain.Payoff(my, op)

		if myPayoff > opPayoff {
			return my
		}

		switch my {
		case domain.Take:
			return domain.Share
		case domain.Share:
			return domain.Hold
		default:
			return domain.Take
		}
	}
}
