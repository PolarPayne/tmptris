package main

import "bytes"

type buttons struct {
	up    bool
	right bool
	down  bool
	left  bool
	cw    bool
	ccw   bool
}

func (b *buttons) reset() {
	b.up = false
	b.right = false
	b.down = false
	b.left = false
	b.cw = false
	b.ccw = false
}

func (b buttons) String() string {
	var buffer bytes.Buffer

	if b.up {
		buffer.WriteString("|UP")
	}
	if b.right {
		buffer.WriteString("|RIGHT")
	}
	if b.down {
		buffer.WriteString("|DOWN")
	}
	if b.left {
		buffer.WriteString("|LEFT")
	}
	if b.cw {
		buffer.WriteString("|CW")
	}
	if b.ccw {
		buffer.WriteString("|CCW")
	}

	st := buffer.String()
	if len(st) > 0 {
		st = st[1:]
	}
	return "buttons[" + st + "]"
}
