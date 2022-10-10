package kitty

type Instruction rune

func (i Instruction) String() string {
	return string(rune(i))
}
