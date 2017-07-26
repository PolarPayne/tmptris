package main

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

type piece struct {
	startState int
	startX     int
	startY     int
	states     []pieceState
}

func (p piece) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(
		"piece[startState=%v startX=%v startY=%v]\n",
		p.startState, p.startX, p.startY))

	for i := range p.states {
		b.WriteString(fmt.Sprintf("state[%v]\n", i))
		b.WriteString(p.states[i].String())
		if i != len(p.states)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

/*
Format of tetrominoes.data:
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
func parseTetriminoes(filename string) map[string]piece {
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

	pieces := make(map[string]piece)
	amountOfPieces := nextInt()

	for i := 0; i < amountOfPieces; i++ {
		name := nextString()
		thePiece := piece{}
		thePiece.startState = nextInt()
		thePiece.startX = nextInt()
		thePiece.startY = nextInt()
		thePiece.states = make([]pieceState, nextInt())

		stateWidth := nextInt()
		stateHeight := nextInt()

		for j := range thePiece.states {
			for y := 0; y < stateHeight; y++ {
				line := nextString()
				for x := 0; x < stateWidth; x++ {
					if line[x] == '1' {
						thePiece.states[j][x][y] = true
					}
				}
			}
		}

		pieces[name] = thePiece

		//fmt.Printf("%v", pieces[i])
	}
	for i := range pieces {
		fmt.Println(pieces[i])
		fmt.Println()
	}
	return pieces
}
