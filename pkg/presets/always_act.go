package presets

// always share
func AlwaysShare() *domain.Strategy {
	return &domain.Strategy{
		Neutral: domain.RuleValue{
			Fix: domain.Share,
		},
		Trigger: nil,
		State:   make(map[string]int),
	}
}

// always hold

// always take
