package day21

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := NewComputer()

	res, err := testDay.Part1(days.Input{
		"A B C D (contains 1, 2)",
		"A E F G (contains 1)",
		"C F (contains 3)",
		"A C G (contains 2)",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "5" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := NewComputer()

	res, err := testDay.Part2(days.Input{
		"A B C D (contains 1, 2)",
		"A E F G (contains 1)",
		"C F (contains 3)",
		"A C G (contains 2)",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "A,C,F" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestParseRecipe(t *testing.T) {
	expectedIngredients := []string{
		"mxmxvkd", "kfcds", "sqjhc", "nhms",
	}
	expectedAllergens := []string{
		"dairy", "fish",
	}
	ingredients, allergens := NewComputer().parseRecipe("mxmxvkd kfcds sqjhc nhms (contains dairy, fish)")

	if !reflect.DeepEqual(expectedIngredients, ingredients) {
		t.Fatalf("Wrong ingredients: %v", ingredients)
	}

	if !reflect.DeepEqual(expectedAllergens, allergens) {
		t.Fatalf("Wrong allergens: %v", allergens)
	}
}

func TestBuildAllergenCandidates(t *testing.T) {
	expectedState := map[string]map[string]int{
		"1": map[string]int{
			"A": 1,
			"B": 1,
			"C": 1,
			"D": 1,
		},
		"2": map[string]int{
			"A": 1,
			"B": 1,
			"C": 1,
			"D": 1},
	}

	testDay := NewComputer()

	testDay.addAllergenCandidates("A B C D (contains 1, 2)")

	if !reflect.DeepEqual(testDay.allergenCandidates, expectedState) {
		t.Fatalf("Invalid candidates %v", testDay.allergenCandidates)
	}
}

func TestBuildAllergenCandidates_TwoRounds(t *testing.T) {
	expectedState := map[string]map[string]int{
		"1": map[string]int{
			"A": 2,
			"B": 1,
			"C": 1,
			"D": 1,
			"E": 1,
			"F": 1,
			"G": 1,
		},
		"2": map[string]int{
			"A": 1,
			"B": 1,
			"C": 1,
			"D": 1},
	}

	maxCounts := map[string]int{
		"1": 2,
		"2": 1,
	}

	testDay := NewComputer()

	testDay.addAllergenCandidates("A B C D (contains 1, 2)")
	testDay.addAllergenCandidates("A E F G (contains 1)")

	if !reflect.DeepEqual(testDay.allergenCandidates, expectedState) {
		t.Fatalf("Invalid candidates %v", testDay.allergenCandidates)
	}

	if !reflect.DeepEqual(testDay.allergensMaxIngredientCount, maxCounts) {
		t.Fatalf("Invalid max counts %v", testDay.allergensMaxIngredientCount)
	}
}

func TestIngredientAppearancesAfterParsinceTwoRecipes(t *testing.T) {
	testDay := NewComputer()

	testDay.parseRecipe("A B C D (contains 1, 2)")
	testDay.parseRecipe("A E F G (contains 1)")

	if len(testDay.ingredients) != 8 {
		t.Fatalf("Invalid ingredient appearances %v", testDay.ingredients)
	}
}
