package day21

import (
	"fmt"
	"sort"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 21
type Computer struct {
	allergenCandidates          map[string]map[string]int
	allergensIngredients        map[string]bool
	allergensMaxIngredientCount map[string]int
	ingredients                 []string
	allergens                   map[string]string
}

// Part1 of Day 21
func (d *Computer) Part1(input days.Input) (days.Result, error) {

	for _, recipe := range input {
		d.addAllergenCandidates(recipe)
	}

	d.findAllergensIngredients()

	return days.Result(fmt.Sprint(d.countNonAllergenAppearances())), nil
}

// Part2 of Day 21
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	d.Part1(input)

	return days.Result(fmt.Sprint(d.refineAllergens())), nil
}

func (d *Computer) refineAllergens() string {
	fmt.Printf("BEFORE REFINEMENT:\n%v\n\n", d.allergenCandidates)

	for len(d.allergenCandidates) != len(d.allergens) {
		for allergen, candidates := range d.allergenCandidates {
			if len(candidates) == 1 {
				for candidate := range candidates {
					for a := range d.allergenCandidates {
						if allergen == a {
							d.allergens[a] = candidate
							continue
						}

						delete(d.allergenCandidates[a], candidate)
					}
				}

			}
		}
	}

	fmt.Printf("AFTER REFINEMENT:\n%v\n\n", d.allergenCandidates)

	sortedAllergens := []string{}

	for allergen := range d.allergens {
		sortedAllergens = append(sortedAllergens, allergen)
	}

	sort.StringSlice(sortedAllergens).Sort()

	res := []string{}
	for _, allergen := range sortedAllergens {
		res = append(res, d.allergens[allergen])
	}

	return strings.Join(res, ",")
}

// NewComputer does stuff
func NewComputer() *Computer {
	return &Computer{
		allergenCandidates:          map[string]map[string]int{},
		ingredients:                 []string{},
		allergensMaxIngredientCount: map[string]int{},
		allergensIngredients:        map[string]bool{},
		allergens:                   map[string]string{},
	}
}

func (d *Computer) countNonAllergenAppearances() int {
	count := 0
	for _, appearance := range d.ingredients {
		if !d.allergensIngredients[appearance] {
			count++
		}
	}
	return count
}

func (d *Computer) findAllergensIngredients() {
	for allergen, candidates := range d.allergenCandidates {
		for candidate, count := range candidates {
			if d.allergensMaxIngredientCount[allergen] == count {
				d.allergensIngredients[candidate] = true
			} else {
				delete(d.allergenCandidates[allergen], candidate)
			}
		}
	}
}

// mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
func (d *Computer) parseRecipe(recipe string) ([]string, []string) {
	halves := strings.Split(recipe, " (contains ")

	ingredients := strings.Split(halves[0], " ")
	allergens := strings.Split(strings.ReplaceAll(halves[1], ")", ""), ", ")

	for _, ingredient := range ingredients {
		d.ingredients = append(d.ingredients, ingredient)
	}

	return ingredients, allergens
}

func (d *Computer) addAllergenCandidates(recipe string) {
	ingredients, allergens := d.parseRecipe(recipe)

	for _, allergen := range allergens {
		if _, ok := d.allergenCandidates[allergen]; !ok {
			d.allergenCandidates[allergen] = map[string]int{}
			d.allergensMaxIngredientCount[allergen] = 0
		}

		for _, ingredient := range ingredients {
			if _, ok := d.allergenCandidates[allergen][ingredient]; !ok {
				d.allergenCandidates[allergen][ingredient] = 0
			}

			d.allergenCandidates[allergen][ingredient]++

			if d.allergenCandidates[allergen][ingredient] > d.allergensMaxIngredientCount[allergen] {
				d.allergensMaxIngredientCount[allergen] = d.allergenCandidates[allergen][ingredient]
			}
		}
	}
}
