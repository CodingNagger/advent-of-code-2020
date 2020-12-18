package day18

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"2 * 3 + (4 * 5)",
		"5 + (8 * 3 + 9 + 3 * 4 * 3)",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "463" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_OneComplexLine(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "12240" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_OneMoreComplexLine(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "13632" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	tests := [][]string{
		[]string{"1 + (2 * 3) + (4 * (5 + 6))", "51"},
		[]string{"2 * 3 + (4 * 5)", "46"},
		[]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)", "1445"},
		[]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", "669060"},
		[]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", "23340"},
	}

	for _, test := range tests {
		res, err := testDay.Part2(days.Input{
			test[0],
		})

		if err != nil {
			t.Fatalf(err.Error())
		}

		if res != days.Result(test[1]) {
			t.Fatalf("Wrong result: %s instead of %s", res, test[1])
		}
	}

}

func TestProcessInnerMostParentheses(t *testing.T) {
	res := processInnerMostParentheses("(1 + 2)")

	if res != "3" {
		t.Fatalf("Didn't process %s", res)
	}
}

func TestProcessInnerMostParentheses_InnerInner(t *testing.T) {
	res := processInnerMostParentheses("(2 + (1 + 2))")

	if res != "5" {
		t.Fatalf("Didn't process %s", res)
	}
}

func TestProcessInnerMostParentheses_Product(t *testing.T) {
	res := processInnerMostParentheses("(2 + (1 * 2))")

	if res != "4" {
		t.Fatalf("Didn't process %s", res)
	}
}

func TestParseAdvancedMathExpression(t *testing.T) {
	expectedResult := productExp{
		constantExp(4),
		sumExp{
			constantExp(2),
			constantExp(3),
		},
	}
	result := parseAdvancedMathExpression("4 * 2 + 3")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseAdvancedMathExpression_TripleProduct(t *testing.T) {
	expectedResult := productExp{
		constantExp(4),
		productExp{
			constantExp(2),
			constantExp(3),
		},
	}
	result := parseAdvancedMathExpression("4 * 2 * 3")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseAdvancedMathExpression_ButMore(t *testing.T) {
	expectedResult := productExp{
		constantExp(2),
		sumExp{
			constantExp(3),
			productExp{
				constantExp(4),
				constantExp(5),
			},
		},
	}
	result := parseAdvancedMathExpression("2 * 3 + (4 * 5)")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseAdvancedMathExpression_EvenMore(t *testing.T) {
	expectedResult := productExp{
		productExp{
			constantExp(5),
			constantExp(9),
		},
		productExp{
			productExp{
				constantExp(7),
				constantExp(3),
			},
			productExp{
				sumExp{
					constantExp(3),
					constantExp(9),
				},
				constantExp(3),
			},
		},
	}
	// result := parseAdvancedMathExpression("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	result := parseAdvancedMathExpression("5 * 9 * (7 * 3 * 3 + 9 * 3)")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v instead of %v", result, expectedResult)
	}
}

func TestParseAdvancedMathExpressionWithParentheses(t *testing.T) {
	expectedResult := productExp{
		constantExp(1),
		sumExp{
			productExp{
				constantExp(4),
				constantExp(2),
			},
			constantExp(3),
		},
	}
	result := parseAdvancedMathExpression("(4 * 2) + 3")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseExpression_TwoOperations(t *testing.T) {
	expectedResult := productExp{
		sumExp{
			constantExp(1),
			constantExp(2),
		},
		constantExp(3),
	}
	result := parseExpression("1 + 2 * 3")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseExpression_Parentheses(t *testing.T) {
	expectedResult := sumExp{
		constantExp(1),
		productExp{
			constantExp(2),
			constantExp(3),
		},
	}
	result := parseExpression("1 + (2 * 3)")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseExpression_SimpleSum(t *testing.T) {
	expectedResult := sumExp{
		constantExp(1),
		constantExp(2),
	}
	result := parseExpression("1 + 2")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseExpression_SimpleProduct(t *testing.T) {
	expectedResult := productExp{
		constantExp(1),
		constantExp(2),
	}
	result := parseExpression("1 * 2")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}
