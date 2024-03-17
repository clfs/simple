package epd

// An Op is an EPD operation.
type Op interface {
	UnmarshalOp(text []byte) error
}

// Known opcodes.
const (
	OpcodeFMVN = "fmvn" // full move number
	OpcodeHMVC = "hmvc" // half move clock
)

// FMVN represents the "fmvn" operation.
type FMVN struct {
	FullMoveNumber int
}

// HMVC represents the "hmvc" operation.
type HMVC struct {
	HalfMoveClock int
}

// func Decode2(s string) (core.Position, []any, error) {
// 	fields := strings.SplitN(s, " ", 5)

// 	if n := len(fields); n < 4 {
// 		return core.Position{}, nil, fmt.Errorf("too few fields: %d", n)
// 	}

// 	pseudoFEN := fmt.Sprintf("%s 0 1", strings.Join(fields[:4], " "))

// 	p, err := fen.Decode(pseudoFEN)
// 	if err != nil {
// 		return core.Position{}, nil, err
// 	}

// 	// Return early if there are no operations.
// 	if len(fields) == 4 {
// 		return p, nil, nil
// 	}

// 	var ops []any

// 	for _, rawOp := range strings.Split(fields[4], ";") {
// 		opcode, operands, _ := strings.Cut(rawOp, " ")
// 		switch {
// 		case opcode == OpcodeFMVN:
// 			n, err := strconv.Atoi(operands)
// 			if err != nil {
// 				return core.Position{}, nil, fmt.Errorf("invalid full move number: %s", operands)
// 			}
// 			ops = append(ops, FMVN{n})
// 		case opcode == OpcodeHMVC:
// 			n, err := strconv.Atoi(operands)
// 			if err != nil {
// 				return core.Position{}, nil, fmt.Errorf("invalid half move clock: %s", operands)
// 			}
// 			ops = append(ops, HMVC{n})
// 		default:
// 			return core.Position{}, nil, fmt.Errorf("unknown opcode: %s", opcode)
// 		}
// 	}
// }
