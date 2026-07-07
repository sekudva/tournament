package agents

import (
	"github.com/sekudva/strategika/internal/domain"
	"github.com/sekudva/strategika/presets/mod"
	"github.com/sekudva/strategika/presets/strategies"
)

// ========== 13. ORIGINAL TOURNAMENT ==========

// GoByMajority - K31R by PAULA GAIL GRISELL.
func GoByMajority() *domain.Agent {
	return &domain.Agent{
		Name:      "GoByMajority",
		ID:        RequestID(130),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(1, mod.GoByMajorityMod())},
	}
}

// Grudger - K34R by JAMES W. FRIEDMAN.
func Grudger() *domain.Agent {
	return &domain.Agent{
		Name:     "Grudger",
		ID:       RequestID(131),
		Strategy: strategies.Grudger(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}

// WMAdams - K44R by WM. ADAMS.
func WMAdams() *domain.Agent {
	return &domain.Agent{
		Name:      "WMAdams",
		ID:        RequestID(132),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(2, mod.AdamsMod())},
	}
}

// Eatherley - K46R by GRAHAM J. EATHERLEY
func Eatherley() *domain.Agent {
	return &domain.Agent{
		Name:      "Eatherley",
		ID:        RequestID(133),
		Strategy:  strategies.TitForTat(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(1, mod.EatherleyMod())},
	}
}

// Cave - K49R by ROB CAVE.
func Cave() *domain.Agent {
	return &domain.Agent{
		Name:      "Cave",
		ID:        RequestID(134),
		Strategy:  strategies.Random(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.CaveMod()},
	}
}

// Pavlov — Win-Stay, Lose-Shift
func Pavlov() *domain.Agent {
	return &domain.Agent{
		Name:      "Pavlov",
		ID:        RequestID(135),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(1, mod.Pavlov())},
	}
}

// Tullock — первые 11 ходов Share, потом вероятностная
func Tullock() *domain.Agent {
	return &domain.Agent{
		Name:      "Tullock",
		ID:        RequestID(136),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(11, mod.TullockMod())},
	}
}

// Champion — 10 Share, 15 TFT, потом условный
func Champion() *domain.Agent {
	return &domain.Agent{
		Name:      "Champion",
		ID:        RequestID(137),
		Strategy:  strategies.TitForTatNEW(), // TFT
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.ChampionMod()},
	}
}

func Leyvraz() *domain.Agent {
	return &domain.Agent{
		Name:      "Leyvraz",
		ID:        RequestID(138),
		Strategy:  strategies.AlwaysShare(),
		Memory:    domain.NewMemory(),
		Score:     0,
		Modifiers: []domain.Modifier{mod.Sleep(1, mod.LeyvrazMod())},
	}
}

func GroFman() *domain.Agent {
	return &domain.Agent{
		Name:     "GroFman",
		ID:       RequestID(139),
		Strategy: strategies.GroFman(),
		Memory:   domain.NewMemory(),
		Score:    0,
	}
}
