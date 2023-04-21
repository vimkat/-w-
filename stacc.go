package kitty

import "fmt"

type Stacc struct {
	data []float64
	sp   uint32
}

func (s *Stacc) String() string {
	return fmt.Sprint(s.data)
}

func (s *Stacc) Copy() *Stacc {
	return s.FirstN(s.Height())
}

func (s *Stacc) Peek() float64 {
	return s.data[s.sp]
}

func (s *Stacc) Pop() float64 {
	ret := s.data[s.sp]
	s.sp--
	return ret
}

func (s *Stacc) Push(v ...float64) *Stacc {
	s.data = append(s.data[:s.Height()], v...)
	s.sp += uint32(len(v))
	return s
}

func (s *Stacc) Height() int {
	return int(s.sp) + 1
}

func (s *Stacc) FirstN(n int) *Stacc {
	ret := &Stacc{sp: s.sp}
	copy(ret.data, s.data[:n])
	return ret
}

func (s *Stacc) Append(other *Stacc) {
	s.data = append(s.data[:s.Height()], other.data[:other.Height()]...)
	s.sp += other.sp
}

func (s *Stacc) Reverse() *Stacc {
	for i := uint32(0); i < s.sp/2; i++ {
		s.data[i], s.data[s.sp-i] = s.data[s.sp-i], s.data[i]
	}

	return s
}

func (s *Stacc) Swap() {
	temp := s.Peek()
	s.data[s.sp] = s.data[0]
	s.data[0] = temp
}

func (s *Stacc) Rotate() {
	last := s.Peek()
	s.data[s.sp] = s.data[len(s.data)-2]
	s.data[len(s.data)-2] = s.data[len(s.data)-3]
	s.data[len(s.data)-3] = last
}

func (s *Stacc) ShiftLeft() *Stacc {
	first := s.data[0]
	s.data = s.data[1:s.Height()]
	s.data[s.sp] = first
	return s
}

func (s *Stacc) ShiftRight() *Stacc {
	last := s.Peek()
	s.data = append([]float64{last}, s.data[:s.sp-1]...)
	return s
}
