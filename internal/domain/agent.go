package domain

type Agent struct {
	ID        AgID
	Memory    *Memory
	Strategy  *Strategy
	Modifiers []Modifier
	Score     int
}

func NewAgent(strat *Strategy, id AgID) *Agent {
	return &Agent{
		ID:       id,
		Strategy: strat,
		Memory:   NewMemory(),
		Score:    0,
	}
}

func (a *Agent) Decide(opID AgID, round int) Act {
	core := a.Strategy.CoreDecision(a.Memory, opID)
	ctx := ModContext{
		Memory:   a.Memory,
		OpID:     opID,
		Round:    round,
		Strategy: a.Strategy,
	}
	for _, mod := range a.Modifiers {
		core = mod(core, ctx)
	}
	return core
}
