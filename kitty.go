package kitty

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
)

type Kitty struct {
	isDead       bool // has stopped execution because of ; or error
	isObserved   bool // debug mode
	laserPointer vec2
	direction    vec2
	litterbox    [][]rune      // code split into lines
	cbbox        *CardboardBox // current stack of values + register[s]
	yarnMode     bool
	out          chan string
	in           <-chan float64
	//cbboxes      []CardboardBox
}

func New(src string) (k *Kitty) {
	k = &Kitty{
		litterbox: readSrc(src),
		direction: vec2{1, 0},
		out:       make(chan string),
		in:        make(chan float64),
	}

	return
}

func (k * Kitty) Out() <-chan string {
	return k.out
}

func (k * Kitty) In(i chan float64) *Kitty {
	k.in = i
	return k
}

func readSrc(src string) [][]rune {
	lines := strings.Split(src, "\n")
	ret := make([][]rune, len(lines))
	for i, line := range lines {
		ret[i] = []rune(line)
	}
	return ret
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
	if k.isDead {
		return k
	}

	instr := k.Fetch()
	k.Execute(instr)
	k.Move()

	return k
}

func (k *Kitty) Move() *Kitty {
		k.laserPointer.x += k.direction.x
		k.laserPointer.y += k.direction.y
	return k
}

func (k *Kitty) IsObserved() bool {
	return k.isObserved
}

func (k *Kitty) CurrentStacc() *Stacc {
	return k.cbbox.Stacc.Copy()
}

func (k *Kitty) Fetch() Instruction {
	return Instruction(k.litterbox[k.laserPointer.y][k.laserPointer.x])
}

func (k *Kitty) Execute(instr Instruction) *Kitty {
	switch instr {
	case NOP: break
	case ADD:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.cbbox.Stacc.Push(x + y)
	case SUB:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.cbbox.Stacc.Push(x - y)
	case MUL:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.cbbox.Stacc.Push(x * y)
	case DIV:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.cbbox.Stacc.Push(x / y)
	case REMAINDER:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.cbbox.Stacc.Push(math.Remainder(x, y))
	case DIE:
		k.isDead = true
	case ROTATE:
		k.cbbox.Stacc.Rotate()
	case JMP:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		k.laserPointer.x = int64(x)
		k.laserPointer.y = int64(y)
	case SKIP:
		k.Move()
	case RIGHT:
		k.direction.x = 1
		k.direction.y = 0
	case LEFT:
		k.direction.x = -1
		k.direction.y = 0
	case UP:
		k.direction.x = 0
		k.direction.y = -1
	case DOWN:
		k.direction.x = 0
		k.direction.y = 1
	case RANDOM:
		k.direction.x = int64(rand.Intn(2) - 1)
		k.direction.y = int64(rand.Intn(2) - 1)
	case EQUAL:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		if x == y {
			k.cbbox.Stacc.Push(1)
		} else {
			k.cbbox.Stacc.Push(0)
		}
	case GREATER:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		if x > y {
			k.cbbox.Stacc.Push(1)
		} else {
			k.cbbox.Stacc.Push(0)
		}
	case LESS:
		y := k.cbbox.Stacc.Pop()
		x := k.cbbox.Stacc.Pop()
		if x < y {
			k.cbbox.Stacc.Push(1)
		} else {
			k.cbbox.Stacc.Push(0)
		}
	case VMIRROR:
		k.direction.x = -k.direction.x
	case HMIRROR:
		k.direction.y = -k.direction.y
	case MIRROR:
		k.direction.x = -k.direction.x
		k.direction.y = -k.direction.y
	case HUH:
		v := k.cbbox.Stacc.Pop()
		if v == 0 {
			k.Move()
		}
	case OPEN:
		hiss()
	case CLOSE:
		hiss()
	case SHIFT_LEFT:
		k.cbbox.Stacc.ShiftLeft()
	case SHIFT_RIGHT:
		k.cbbox.Stacc.ShiftRight()
	case DUP:
		k.cbbox.Stacc.Push(k.cbbox.Stacc.Peek())
	case DISCARD:
		k.cbbox.Stacc.Pop()
	case SWAP:
		k.cbbox.Stacc.Swap()
	case LENGTH:
		k.cbbox.Stacc.Push(float64(k.cbbox.Stacc.Height()))
	case PUT:
		y := int(k.cbbox.Stacc.Pop())
		x := int(k.cbbox.Stacc.Pop())
		v := rune(k.cbbox.Stacc.Pop())
		k.litterbox[y][x] = v
	case GET:
		y := int(k.cbbox.Stacc.Pop())
		x := int(k.cbbox.Stacc.Pop())
		k.cbbox.Stacc.Push(float64(k.litterbox[y][x]))
	case YARN:
		k.yarnMode = !k.yarnMode
	case OUTCHR:
		k.out<-string(rune(k.cbbox.Stacc.Pop()))
	case OUTNUM:
		k.out<-fmt.Sprintf("%f", k.cbbox.Stacc.Pop())
	case INPUT:
		k.cbbox.Stacc.Push(<-k.in)
	default:
		chr := rune(instr)
		if isHex(chr) {
			k.cbbox.Stacc.Push(fromHex(chr))
			break
		}
		if k.yarnMode {
			k.cbbox.Stacc.Push(float64(chr))
			break
		}
		hiss()
	}

	return k
}

func isHex(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'A' && r <= 'F')
}

func fromHex(r rune) float64 {
	if r >= '0' && r <= '9' {
		return float64(r - '0')
	}
	if r >= 'A' && r <= 'F' {
		return float64(r - 'A' + 10)
	}
	return -1
}

func hiss() {
	fmt.Fprintf(os.Stderr, "*HISS!*\n")
	os.Exit(1)
}
