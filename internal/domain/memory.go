package domain

type Round struct {
	N         int
	PartnerID AgID
	MyAct     Act
	OpAct     Act
}

type Memory struct {
	History map[AgID][]Round
}

func NewMemory() *Memory {
	return &Memory{
		History: make(map[AgID][]Round),
	}
}

// Запись хода в память агента
func (m *Memory) Record(round int, partnerID AgID, my, their Act) {
	m.History[partnerID] = append(m.History[partnerID], Round{
		N:         round,
		PartnerID: partnerID,
		MyAct:     my,
		OpAct:     their,
	})
}

// Возвращает последний ход ПРОТИВНИКА
// Если это первый ход в истории то возвращает -1
func (m *Memory) OpLastAct(partnerID AgID) Act {
	history, ok := m.History[partnerID]
	if !ok || len(history) == 0 {
		return NoAct
	}
	return history[len(history)-1].OpAct
}

// Возвращает МОЙ последний ход
// Если это первый ход в истории то возвращает -1
func (m *Memory) MyLastAct(partnerID AgID) Act {
	history, ok := m.History[partnerID]
	if !ok || len(history) == 0 {
		return NoAct
	}
	return history[len(history)-1].MyAct
}

// Счетчик триггера для стратегий
func (m *Memory) CountTrigger(partnerID AgID, trigger Act) int {
	history, here := m.History[partnerID]
	if !here {
		return 0
	}

	count := 0
	for _, round := range history {
		if round.OpAct == trigger {
			count++
		}
	}

	return count
}
