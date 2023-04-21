package kitty

type Instruction rune

const (
	NOP         Instruction = ' '
	ADD         Instruction = '+'
	SUB         Instruction = '-'
	MUL         Instruction = '*'
	DIV         Instruction = '/'
	REMAINDER   Instruction = '%'
	DIE         Instruction = ';'
	ROTATE      Instruction = '@'
	JMP         Instruction = '.'
	SKIP        Instruction = '!'
	RIGHT       Instruction = '→'
	LEFT        Instruction = '←'
	UP          Instruction = '↑'
	DOWN        Instruction = '↓'
	RANDOM      Instruction = 'x'
	EQUAL       Instruction = '='
	GREATER     Instruction = '>'
	LESS        Instruction = '<'
	VMIRROR     Instruction = '|'
	HMIRROR     Instruction = '_'
	MIRROR      Instruction = '#'
	HUH         Instruction = '?'
	OPEN        Instruction = '['
	CLOSE       Instruction = ']'
	SHIFT_LEFT  Instruction = '{'
	SHIFT_RIGHT Instruction = '}'
	DUP         Instruction = ':'
	DISCARD     Instruction = '~'
	SWAP        Instruction = '$'
	LENGTH      Instruction = 'l'
	PUT         Instruction = 'p'
	GET         Instruction = 'g'
	YARN        Instruction = '"'
	OUTCHR      Instruction = 'o'
	OUTNUM      Instruction = 'n'
	INPUT       Instruction = 'i'
)

func (i Instruction) String() string {
	return string(rune(i))
}
