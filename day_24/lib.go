package day_24

import (
	"strconv"
	"strings"
)

type wire struct {
	name   string
	inputs []*wire
	gate   string
	value  int
}

func (w *wire) resolve() int { // can I do this recursively? I will need some kind of return for that. Return the value if not -1?
	if w.value != -1 {
		return w.value
	}

	switch w.gate {
	case "AND":
		w.value = w.inputs[0].resolve() & w.inputs[1].resolve()
	case "OR":
		w.value = w.inputs[0].resolve() | w.inputs[1].resolve()
	case "XOR":
		w.value = w.inputs[0].resolve() ^ w.inputs[1].resolve()
	}
	return w.value
}

func zVals(ws map[string]*wire) int {
	var out int

	for _, w := range ws {
		if w.name[0] == 'z' {
			power, err := strconv.Atoi(w.name[1:])
			if err != nil {
				panic("Oh boy!")
			}

			out += pow(2, power) * w.resolve()
		}
	}

	return out
}

func pow(b, e int) int {
	out := 1

	for i := 0; i < e; i++ {
		out *= b
	}

	return out
}

func parse(s string) map[string]*wire {
	out := make(map[string]*wire)
	blocks := strings.Split(strings.TrimSpace(s), "\n\n")

	// The first block contains initial values
	for _, line := range strings.Split(blocks[0], "\n") {
		fields := strings.Fields(line)

		name := fields[0][:len(fields[0])-1]
		val, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("Big problem")
		}

		out[name] = &wire{
			name:   name,
			inputs: []*wire{},
			gate:   "",
			value:  val,
		}
	}

	// The second block contains the gate setup
	for _, line := range strings.Split(blocks[1], "\n") {
		fields := strings.Fields(line)

		if _, present := out[fields[0]]; !present {
			out[fields[0]] = &wire{
				name:   fields[0],
				inputs: []*wire{},
				gate:   "",
				value:  -1,
			}
		}

		if _, present := out[fields[2]]; !present {
			out[fields[2]] = &wire{
				name:   fields[2],
				inputs: []*wire{},
				gate:   "",
				value:  -1,
			}
		}

		if w, present := out[fields[4]]; present {
			w.inputs = []*wire{out[fields[0]], out[fields[2]]}
			w.gate = fields[1]
		} else {
			out[fields[4]] = &wire{
				name:   fields[4],
				inputs: []*wire{out[fields[0]], out[fields[2]]},
				gate:   fields[1],
				value:  -1,
			}
		}
	}

	return out
}
