package eval

import (
	"testing"
	"fmt"
	"math"
)

// An Expr is an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in this environment env
	Eval(env Env) float64
}

// A Var identifies a variable, eg. x
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// A literal is a numeric constant, eg. 3.141
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// A unary represents a unary operator expression, e.g -x
type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("Unsupported unary operator: %q", u.op))
}

// A unary represents a binary operator expression, e.g x+y
type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("Unsupported Binary operator: %q", b.op))
}

// A call represents a function call expression e.g sin(x)
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64


// TestEval tests the expression evaluator apparatus above
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string

	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
	}
}