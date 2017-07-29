package board

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const (
	pieceStateWidth  = 4
	pieceStateHeight = 4
)

type pieceState [pieceStateWidth][pieceStateWidth]bool

func (p pieceState) String() string {
	var b bytes.Buffer

	for y := 0; y < pieceStateHeight; y++ {
		for x := 0; x < pieceStateWidth; x++ {
			if p[x][y] {
				b.WriteString("O")
			} else {
				b.WriteString("-")
			}
		}
		if y != pieceStateHeight-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

type Piece struct {
	StartState int
	StartX     int
	StartY     int
	States     []pieceState
}

func (p Piece) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(
		"piece[startState=%v startX=%v startY=%v]\n",
		p.StartState, p.StartX, p.StartY))

	for i := range p.States {
		b.WriteString(fmt.Sprintf("state[%v]\n", i))
		b.WriteString(p.States[i].String())
		if i != len(p.States)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

type Pieces map[string]Piece

/*
ParseFile reads piece information from a file.

Format of the file MUST be like:
	<pieces>
	[for _ in range(pieces)]
	<name>
	<startState> <startX> <startY>
	<states> <stateWidth> <stateHeight>
	[for _ in range(states)]

	[for _ in range(stateHeight)]
	[for _ in range(stateWidth)](0|1)[endfor]
	[endfor]

	[endfor]
	[endfor]
*/
func ParseFile(filename string) Pieces {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open " + filename)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	nextString := func() string {
		if scanner.Scan() {
			return scanner.Text()
		}
		panic("Badly formatted piece data")
	}
	nextInt := func() int {
		s := nextString()
		out, err := strconv.Atoi(s)
		if err != nil {
			panic("Badly formatted number in piece data")
		}
		return out
	}

	pieces := make(Pieces)
	amountOfPieces := nextInt()

	for i := 0; i < amountOfPieces; i++ {
		name := nextString()
		thePiece := Piece{}
		thePiece.StartState = nextInt()
		thePiece.StartX = nextInt()
		thePiece.StartY = nextInt()
		thePiece.States = make([]pieceState, nextInt())

		stateWidth := nextInt()
		stateHeight := nextInt()

		for j := range thePiece.States {
			for y := 0; y < stateHeight; y++ {
				line := nextString()
				for x := 0; x < stateWidth; x++ {
					if line[x] == 'o' {
						thePiece.States[j][x][y] = true
					}
				}
			}
		}

		pieces[name] = thePiece
	}
	return pieces
}
