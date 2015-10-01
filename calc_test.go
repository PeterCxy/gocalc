package calc

import "fmt"
import "testing"
import "strings"
import "math"

func TestCalc(t *testing.T) {
	test(t, "1 + 2", 3)
	test(t, "2 - 4", -2)
	test(t, "2 * 4", 8)
	test(t, "2 / 4", 0.5)
	test(t, "6 % 4", 2)
	test(t, "1.1 * 2.2", 2.42)
	test(t, "1 - (4 + 3)", -6)
	test(t, "- (4 + 4)", -8)
	test(t, "sqrt(4) + 1", 3)
	test(t, "sqrt(4) + sqrt(9)", 5)
	test(t, "2 ^ (32 - 1) - 1", 2147483647)
	test(t, "Floor(1.2)", 1)
	test(t, "ceil(1.2)", 2)
	test(t, "abs(-2.8)", 2.8)
	test(t, "abs(- 2.8 + 1.1)", 1.7)
	test(t, "log(8, 2)", 3)
	test(t, "ceil(e)", 3)
	test(t, "ln(e)", 1)
	test(t, "sin(pi / 6)", 0.5)
	test(t, "cos(pi / 3)", 0.5)
	test(t, "tan(pi / 4)", 1)
	test(t, "arccos(0.5)", math.Pi / 3)
}

func test(t *testing.T, expr string, expect float64) {
	r, err := Calculate(expr)
	fmt.Printf(strings.Replace(expr, "%", "%%", -1) + " = %f, expect %f\n", r, expect)
	if err != nil {
		fmt.Println(err)
	}
	if !floatEquals(r, expect) {
		t.Fail()
	}
}

var EPSILON float64 = 0.00000001
func floatEquals(a, b float64) bool {
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	} else {
		return false
	}
}
