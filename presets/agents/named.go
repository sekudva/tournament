package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/mod"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 10. NEW LOGIC ==========

// Рассказывает в своих статьях о самых знаменательных событиях!
// И почему-то иногда от скуки решает придумать кричащий заголовок сама...
func Journalist() *domain.Agent {
	return &domain.Agent{
		Name: "Journalist",
		ID:   RequestID(100),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:  domain.Hold,
				Prob: map[domain.Act]float64{domain.Take: 0.05},
			},
		},

		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.JournalistMod()},
	}
}

// best to any TFT strategy
// Exploiter to TFT, author: Boev
func Boev() *domain.Agent {
	return &domain.Agent{
		Name: "Boev-AntiTFT",
		ID:   RequestID(101),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix:  domain.Share,
				Prob: map[domain.Act]float64{domain.Take: 0.3},
			},
		},

		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Exploiter()},
	}
}

func Abuser() *domain.Agent {
	return &domain.Agent{
		Name: "Abuser",
		ID:   RequestID(102),

		Strategy: strategies.Random(),

		Memory: domain.NewMemory(),
		Score:  0,
		Modifiers: []domain.Modifier{
			mod.Exploiter(),
			mod.Sleep(15, mod.SeekBrick()),
			mod.Sleep(30, mod.SeekWeak()),
		},
	}
}

// Envy strategy.
func EvilPavlov() *domain.Agent {
	return &domain.Agent{
		Name:      "EvilPavlov",
		ID:        RequestID(103),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(1, mod.EvilPavlov())},
	}
}

// Oracle strategy.
// Based on Branch Predictor Unit, author: Sycheva
func Sycheva() *domain.Agent {
	return &domain.Agent{
		Name: "Sycheva-Oracle",
		ID:   RequestID(104),

		Strategy: &domain.Strategy{
			Neutral: domain.RuleValue{
				Fix: domain.Share,
			},
		},

		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.BPU()},
	}
}
