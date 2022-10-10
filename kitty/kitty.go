package kitty

import "strings"

type Kitty struct {
	isDead       bool // has stopped execution because of ; or error
	isObserved   bool // debug mode
	laserPointer vec2
	direction    vec2
	litterbox    []string      // code split into lines
	cbbox        *CardboardBox // current stack of values + register[s]
}

func New(src string) (k *Kitty) {
	k = &Kitty{
		litterbox: strings.Split(src, "\n"),
		direction: vec2{1, 0},
	}

	return
}

func (k *Kitty) Observe() *Kitty {
	k.isObserved = true
	return k
}

func (k *Kitty) Unobserve() *Kitty {
	k.isObserved = false
	return k
}

func (k *Kitty) IsDead() bool {
	return k.isDead
}

func (k *Kitty) HuntLaser() *Kitty {

	return k
}

func (k *Kitty) IsObserved() bool {
	return k.isObserved
}

func (k *Kitty) CurrentStacc() *Stacc {
	return k.cbbox.Stacc.Copy()
}

func (k *Kitty) CurrentInstruction() Instruction {
	return Instruction(k.litterbox[k.laserPointer.y][k.laserPointer.x])
}
