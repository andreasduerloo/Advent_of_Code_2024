package day_17

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

type computer struct {
	regA     int
	regB     int
	regC     int
	instrPtr int
	instr    []int
	opcodes  []func(*computer, int)
	output   []int
	halt     bool
}

func parse(s string) *computer {
	blocks := strings.Split(s, "\n\n")

	lines := strings.Split(blocks[0], "\n")

	return &computer{
		regA:     helpers.ReGetInts(lines[0])[0],
		regB:     helpers.ReGetInts(lines[1])[0],
		regC:     helpers.ReGetInts(lines[2])[0],
		instrPtr: 0,
		instr:    helpers.ReGetInts(blocks[1]),
		opcodes: []func(*computer, int){
			adv,
			bxl,
			bst,
			jnz,
			bxc,
			out,
			bdv,
			cdv,
		},
		output: []int{},
		halt:   false,
	}
}

func (c *computer) execute() {
	if c.instrPtr >= len(c.instr) {
		c.write()
		c.halt = true
		return
	}

	c.opcodes[c.instr[c.instrPtr]](c, c.instr[c.instrPtr+1])
}

func (c computer) write() {
	var out string

	for i, o := range c.output {
		out += (strconv.Itoa(o))
		if i < len(c.output)-1 {
			out += ","
		}
	}

	fmt.Println(out)
}

func intPow(b, e int) int {
	out := 1

	for i := 0; i < e; i++ {
		out *= b
	}

	return out
}

func handleCombo(c *computer, i int) int {
	var combo int

	if i >= 0 && i < 4 {
		combo = i
	}

	if i == 4 {
		combo = c.regA
	}

	if i == 5 {
		combo = c.regB
	}

	if i == 6 {
		combo = c.regC
	}

	return combo
}

func adv(c *computer, i int) {
	combo := handleCombo(c, i)
	c.regA = c.regA / intPow(2, combo)
	c.instrPtr += 2
}

func bxl(c *computer, i int) {
	c.regB = c.regB ^ i
	c.instrPtr += 2
}

func bst(c *computer, i int) {
	combo := handleCombo(c, i)
	c.regB = combo % 8
	c.instrPtr += 2
}

func jnz(c *computer, i int) {
	if c.regA == 0 {
		c.instrPtr += 2
		return
	}
	c.instrPtr = i
}

func bxc(c *computer, i int) {
	c.regB = c.regB ^ c.regC
	c.instrPtr += 2
}

func out(c *computer, i int) {
	combo := handleCombo(c, i)
	c.output = append(c.output, combo%8)
	c.instrPtr += 2
}

func bdv(c *computer, i int) {
	combo := handleCombo(c, i)
	c.regB = c.regA / intPow(2, combo)
	c.instrPtr += 2
}

func cdv(c *computer, i int) {
	combo := handleCombo(c, i)
	c.regC = c.regA / intPow(2, combo)
	c.instrPtr += 2
}
