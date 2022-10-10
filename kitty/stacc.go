package kitty

type Stacc struct {
	data []float64
	sp   uint32
}

func (s *Stacc) String() string {
	return "a Stacc"
}

func (s *Stacc) Copy() *Stacc {
	//TODO @ethan: actually copy the thing (since: 2022-10-10)
	return s
}
