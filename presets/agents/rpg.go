package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/mod"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 11. RPG CLASSES ==========

// Always nice to harmless, fights with ones who hurting others
func Paladin() *domain.Agent {
	return &domain.Agent{
		Name: "Paladin",
		ID:   RequestID(110),

		Strategy: strategies.TitForTat(),

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.Warrior(),
			mod.Assistant(),
		},
	}
}

// Loves fighting, never holding back
func Barbarian() *domain.Agent {
	b := domain.MirrorBad
	return &domain.Agent{
		Name: "Barbarian",
		ID:   RequestID(111),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:    domain.Share,
				Mirror: &b,
				Prob:   map[domain.Act]float64{domain.Take: 0.02},
			},
		},

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.Warrior(),
			mod.Brawler(),
		},
	}
}

// Attacks harmless and hide from strong
// Sometimes betray
func Rogue() *domain.Agent {
	d := domain.MirrorDefense
	return &domain.Agent{
		Name: "Rogue",
		ID:   RequestID(112),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:    domain.Hold,
				Mirror: &d,
				Prob: map[domain.Act]float64{
					domain.Take:  0.2,
					domain.Share: 0.1,
				},
			},
		},

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.Quicksand(),
			mod.SeekFreak(),
			mod.WithProbability(0.3, mod.SeekBrick()),
		},
	}
}

// Very forgiving, never take
func Monk() *domain.Agent {
	return &domain.Agent{
		Name: "Monk",
		ID:   RequestID(113),

		Strategy: strategies.Saint(),

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.Assistant(),
			mod.WithProbability(0.5, mod.Mature()),
		},
	}
}

// Posessed, but fighting for his mind
func Warlock() *domain.Agent {
	d := domain.MirrorDirect
	return &domain.Agent{
		Name: "Warlock",
		ID:   RequestID(114),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:    domain.Share,
				Mirror: &d,
				Prob: map[domain.Act]float64{
					domain.Take: 0.2,
					domain.Hold: 0.3,
				},
			},
		},

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.WithProbability(0.2, mod.Mature()),
			mod.WithProbability(0.2, mod.Quicksand()),
		},
	}
}

// Unpredictable strategy with unpredictable principles
func Sorcerer() *domain.Agent {
	return &domain.Agent{
		Name: "Sorcerer",
		ID:   RequestID(115),

		Strategy: strategies.Random(),

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.WithProbability(0.3, mod.Quicksand()),
			mod.WithProbability(0.3, mod.Assistant()),
		},
	}
}
