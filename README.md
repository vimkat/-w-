# ^w^ (pronunciation: kɪtɪ)

Homebase of the `^w^` (pronunciation: kɪtɪ) programming language

`^w^` is a 2D langage base on [><>](https://esolangs.org/wiki/Fish).
That means it ...

- is stacc based (everything on the stacc is a float64 and can be converted at will)
- uses number in hex (but offers literals in any base ... soon™)
- supports null-terminated strings in various functions
- offers a way to call dynamic libraries (macOS and Linux for now ... also soon™)

Also, be careful! The language is still in **EARLY DEVELOPMENT**.
Things might break. Things might change. You might have a bad time. You might hate yourself.
But please remember: you're dealing with a kitty! It needs love and it purrs! <3

---

## Hello World

```kitty
"Hello World"0rP;
```


## Numbers and Math

The following characters are recognized as (base 16) digits where each digit pushes itself onto the stacc.

```kitty
0123456789ABCDEF
```

Your stacc now looks like this:

```
[ 0 1 2 3 4 5 6 7 8 9 A B C D E F ]
                                ^
                                SP (stacc pointer), F is referred to as TOS (top of stacc)
```


### Opurrators

Opurrators pop `y` then `x` and push `x [op] y` bacc onto the stacc.
Here are some basic opurrators you might find useful (and also all that kitty supports):

```kitty
3 4 + ; addition        TOS: 7
3 4 - ; subtraction     TOS: -1
3 4 * ; multiplication  TOS: 12 (or 0xC)
3 4 / ; division        TOS: 0.75
3 4 % ; modulo          TOS: 3
```


## Input and Output

To make your program actually useful it has to be able to receive input for an opurrating cat or give output bacc to said cat.
You can use the following instructions:

- `i` reads a caracter from stdin and pushes it onto the stacc
	- reading a number character (`[0-9]`) will push the actual number
	- reading anything else (e. g. `purr`) will push the unicode representation
- `o` outputs TOS interpreted as a unicode charactater
- `n` outputs TOS as a number


## How to write code

`^w^` code is written on an infinite grid called *The Litterbox* where the top left corner has coordinates `(0,0)`.
Execution follows a laser pointer (LP) starting out at `(0,0)` facing to the right.
In case the LP is moved to a negative coordinate (with `.`) the `^w^` gets mad and shall hiss at the opurrating cat.
Negative *Litterbox* coordiantes may be used with the `g` and `p` instructions to store arbitrary data.

The direction may be changed with the instructions `← ↓ ↑ →`.
If the LP goes outside of the bounds of *The Litterbox* defined by the code, it wraps around in the current direction.
There is only one error message, that being `*HISS!*`

Space characters are ignored, making them suitable to structure your *Litterbox*.
Multiple consecutive spaces are to be treated as a single space, which is also optional, making all three of the following programs equivalent:
```kitty
12+n;
1 2 + n;
1 2   +   n                     ;
```
They all take 5 steps to execute and terminate.


## Flow Control

The following opurrations can be used to make decisions in your code:

- `<`, `>`, `=`, `≤`, `≥` each pop `y`, then `x`, and push a 0 if the condition defined by `x [op] y` is false, and a 1 if it is true
- `?` pops the stacc
	- if a nonzero value is read, the immediately following instruction is executed
	- if a zero is read instead, the following instruction is skipped and the one thereafter is executed instead
- `!` skips the next instruction unconditionally
- `.` pops `y`, then `x`, and moves the LP to `(x, y)`, making that the next instruction to be executed
- `;` allows `^w^` to catch the elusive dot produced by the LP, thus making it curl up and sleep happily ever after


## Other instructions

- `:` duplicates the TOS
- `l` pushes the length of the stacc onto the stacc

`TODO @ethan: List other ><> commands that we also have in ^w^ (since: 2022-10-10)`


## Examples

Here are some examples to make things clearer to understand :P


### Add two numbers

```kitty
ii+n;
```


### Flooring a float

```kitty
ii/:1%-n;
```

---

Made with <3 by two crazy catbois

