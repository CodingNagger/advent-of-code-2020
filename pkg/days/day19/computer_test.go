package day19

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		"4: \"a\"",
		"5: \"b\"",
		"",
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_NoChangesApplied(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"42: 9 14 | 10 1",
		"9: 14 27 | 1 26",
		"10: 23 14 | 28 1",
		"1: \"a\"",
		"11: 42 31",
		"5: 1 14 | 15 1",
		"19: 14 1 | 14 14",
		"12: 24 14 | 19 1",
		"16: 15 1 | 14 14",
		"31: 14 17 | 1 13",
		"6: 14 14 | 1 14",
		"2: 1 24 | 14 4",
		"0: 8 11",
		"13: 14 3 | 1 12",
		"15: 1 | 14",
		"17: 14 2 | 1 7",
		"23: 25 1 | 22 14",
		"28: 16 1",
		"4: 1 1",
		"20: 14 14 | 1 15",
		"3: 5 14 | 16 1",
		"27: 1 6 | 14 18",
		"14: \"b\"",
		"21: 14 1 | 1 14",
		"25: 1 1 | 1 14",
		"22: 14 14",
		"8: 42",
		"26: 14 22 | 1 20",
		"18: 15 15",
		"7: 14 5 | 1 21",
		"24: 14 1",
		"",
		"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaaaabbaaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"babaaabbbaaabaababbaabababaaab",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "3" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_SmallScopeNoChanges(t *testing.T) {
	testDay := &Computer{}

	// res, err := testDay.Part2(days.Input{
	// 	"0: 1 2 3 4",
	// 	"1: \"a\"",
	// 	"2: \"b\"",
	// 	"3: 1 | 1 3 2",
	// 	"4: 2 | 2 4",
	// 	"",
	// 	"abab",
	// 	"abaab",
	// 	"ababb",
	// })

	res, err := testDay.Part2(days.Input{
		"0: 1 2 5",
		"1: \"a\"",
		"2: \"b\"",
		"3: 1 | 1 4 1",
		"4: 2 | 2 4",
		"5: 3 | 4 5",
		"",
		"aba",
		"ababa",
		"ababbba",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "3" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_SingleMessage(t *testing.T) {
	testDay := &Computer{AddLoops: true}

	res, err := testDay.Part2(days.Input{
		"42: 9 14 | 10 1",
		"9: 14 27 | 1 26",
		"10: 23 14 | 28 1",
		"1: \"a\"",
		"11: 42 31",
		"5: 1 14 | 15 1",
		"19: 14 1 | 14 14",
		"12: 24 14 | 19 1",
		"16: 15 1 | 14 14",
		"31: 14 17 | 1 13",
		"6: 14 14 | 1 14",
		"2: 1 24 | 14 4",
		"0: 8 11",
		"13: 14 3 | 1 12",
		"15: 1 | 14",
		"17: 14 2 | 1 7",
		"23: 25 1 | 22 14",
		"28: 16 1",
		"4: 1 1",
		"20: 14 14 | 1 15",
		"3: 5 14 | 16 1",
		"27: 1 6 | 14 18",
		"14: \"b\"",
		"21: 14 1 | 1 14",
		"25: 1 1 | 1 14",
		"22: 14 14",
		"8: 42",
		"26: 14 22 | 1 20",
		"18: 15 15",
		"7: 14 5 | 1 21",
		"24: 14 1",
		"",
		"bbbababbbbaaaaaaaabbababaaababaabab",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{AddLoops: true}

	res, err := testDay.Part2(days.Input{
		"42: 9 14 | 10 1",
		"9: 14 27 | 1 26",
		"10: 23 14 | 28 1",
		"1: \"a\"",
		"11: 42 31",
		"5: 1 14 | 15 1",
		"19: 14 1 | 14 14",
		"12: 24 14 | 19 1",
		"16: 15 1 | 14 14",
		"31: 14 17 | 1 13",
		"6: 14 14 | 1 14",
		"2: 1 24 | 14 4",
		"0: 8 11",
		"13: 14 3 | 1 12",
		"15: 1 | 14",
		"17: 14 2 | 1 7",
		"23: 25 1 | 22 14",
		"28: 16 1",
		"4: 1 1",
		"20: 14 14 | 1 15",
		"3: 5 14 | 16 1",
		"27: 1 6 | 14 18",
		"14: \"b\"",
		"21: 14 1 | 1 14",
		"25: 1 1 | 1 14",
		"22: 14 14",
		"8: 42",
		"26: 14 22 | 1 20",
		"18: 15 15",
		"7: 14 5 | 1 21",
		"24: 14 1",
		"",
		"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaaaabbaaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"babaaabbbaaabaababbaabababaaab",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "12" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestIsRule(t *testing.T) {
	tests := map[string]bool{
		"0: 4 1 5":     true,
		"3: 4 5 | 5 4": true,
		"4: \"a\"":     true,
		"":             false,
		"bababa":       false,
	}

	for test, expectedResult := range tests {
		if isRule(test) != expectedResult {
			t.Fatalf("Expected isRule to return %v for %s", expectedResult, test)
		}
	}
}

func TestParseRule_ReturnsTheRightRule(t *testing.T) {
	rule3 := charRule{"a"}
	rule4 := charRule{"b"}

	tests := map[string]interface{}{
		"0: 4 1 3": andRuleReference{ruleReference(4), ruleReference(1), ruleReference(3)},
		"1: 4 4 | 3 3": orRuleReference{
			andRuleReference{ruleReference(4), ruleReference(4)},
			andRuleReference{ruleReference(3), ruleReference(3)},
		},
		"2: 4 3 | 3 4": orRuleReference{
			andRuleReference{ruleReference(4), ruleReference(3)},
			andRuleReference{ruleReference(3), ruleReference(4)},
		},
		"3: \"a\"": rule3,
		"4: \"b\"": rule4,
	}

	for test, expectedResult := range tests {
		_, rule := parseRule(test)
		if !reflect.DeepEqual(rule, expectedResult) {
			t.Fatalf("Expected parseRule(%s) to return %v not %v", test, expectedResult, rule)
		}
	}
}
