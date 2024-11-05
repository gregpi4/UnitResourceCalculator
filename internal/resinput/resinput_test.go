package resinput

import (
	"log"
	"reflect"
	"testing"
)

func deepEqualUnordered(listA []string, listB []string) bool {
	setA := make(map[string]struct{}, len(listA))
	for _, s := range listA {
		setA[s] = struct{}{}
	}
	if len(setA) != len(listB) {
		return false
	}
	for _, item := range listB {
		if _, ok := setA[item]; !ok {
			log.Printf("hello there")
			return false
		}
	}
	return true
}

// TestCalculateA returns number of base atom needed to produce product A
func TestGetProductsName(t *testing.T) {
	want := []string{"bell", "cog", "inserter", "automate"}
	base_inputs := GetProductsName("vanilla")
	if !deepEqualUnordered(want, base_inputs) {
		t.Fatalf(`Calculate("A") = %q, want match for %#q, nil`, base_inputs, want)
	}
}

// TestDesign with 1 product and 1 recipe returns [["", "", "automate", "automate", "automate", "", ""], ["bell_cog", "inserter_cog", "automate_cog", "automate_cog", "automate_cog", "inserter_ironPlate", "bell_ironPlate"], ["", "", "automate", "automate", "automate", "", ""]]
func TestDesignCog(t *testing.T) {
	want := [][]string{
		{"", "", "undergroundBell_ironPlate", "", ""},
		{"undergroundBell_cog", "inserter_cog", "automate_cog", "automate_cog", "automate_cog"},
		{"", "", "automate_cog", "automate_cog", "automate_cog"},
		{"", "", "automate_cog", "automate_cog", "automate_cog"},
		{"", "", "inserter", "", ""},
		{"", "", "undergroundBell_ironPlate", "", ""},
	}
	base_inputs, _ := Design("vanilla", "cog", []string{"ironPlate"})
	if !reflect.DeepEqual(want, base_inputs) {
		t.Fatalf(`Calculate("A") = %q, want match for %#q, nil`, base_inputs, want)
	}
}

// TestDesign with 1 product and 1 recipe returns [["", "", "automate", "automate", "automate", "", ""], ["bell_cog", "inserter_cog", "automate_cog", "automate_cog", "automate_cog", "inserter_ironPlate", "bell_ironPlate"], ["", "", "automate", "automate", "automate", "", ""]]
func TestDesignBell(t *testing.T) {
	want := [][]string{
		{"", "", "undergroundBell_ironPlate_cog", "", ""},
		{"undergroundBell_bell", "inserter_bell", "automate_bell", "automate_bell", "automate_bell"},
		{"", "", "automate_bell", "automate_bell", "automate_bell"},
		{"", "", "automate_bell", "automate_bell", "automate_bell"},
		{"", "", "inserter", "", ""},
		{"", "", "undergroundBell_ironPlate_cog", "", ""},
	}
	base_inputs, _ := Design("vanilla", "bell", []string{"ironPlate", "cog"})
	if !reflect.DeepEqual(want, base_inputs) {
		t.Fatalf(`Calculate("A") = %q, want match for %#q, nil`, base_inputs, want)
	}
}
