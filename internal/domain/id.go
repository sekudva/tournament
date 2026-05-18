package domain

type AgID uint32

type IDGenerator struct {
	lastID AgID
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		lastID: 1,
	}
}

func (g *IDGenerator) Next() AgID {
	g.lastID++
	return g.lastID
}

func (g *IDGenerator) Request(id AgID) AgID {
	if id > g.lastID {
		g.lastID = id
	}
	return id
}
