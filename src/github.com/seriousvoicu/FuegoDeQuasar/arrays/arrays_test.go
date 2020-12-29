package arrays

import (
	"testing"
)

func TestMergeStringArrays(t *testing.T) {
	var testTable = []struct {
		testName string
		arrayA   []string
		arrayB   []string
		spected  string
	}{
		{
			"Elementos no coincidentes",
			[]string{"a", "b", ""},
			[]string{"", "c", "d"},
			"No hay coincidencias (arrays.string_array)",
		},
		{
			"Elementos sin definir al final",
			[]string{"a", "b", "", ""},
			[]string{"", "", "a", "", "c", ""},
			"a b c ",
		},
		{
			"ok 1",
			[]string{"", "", "a", "b", "", "d"},
			[]string{"a", "", "c", ""},
			"a b c d",
		},
		{
			"ok 2",
			[]string{"", "", "a", "b", "", "d"},
			[]string{"", "", "c", ""},
			"a b c d",
		},
	}

	for _, ii := range testTable {
		t.Run(ii.testName, func(t *testing.T) {
			result, state := MergeStringArrays(ii.arrayA, ii.arrayB)

			if !state.IsOk() {
				if state.UserError != "No hay coincidencias (arrays.string_array)" {
					t.Errorf(state.UserError)
					return
				}

				return
			}

			str, state := StringArrayToString(result)

			if !state.IsOk() {
				t.Errorf(state.UserError)
				return
			}

			if str != ii.spected {
				t.Errorf("got %q, want %q", str, ii.spected)
				return
			}

		})
	}

}

func TestArrayToString(t *testing.T) {
	var testTable = []struct {
		testName string
		array    []string
		spected  string
	}{
		{"testA", []string{"a", "b", "c", "d"}, "a b c d"},
		{"testB", []string{"a", "b", "", ""}, "a b  "},
		{"testC", []string{" ", "", "a", "b", "", ""}, "   a b  "},
	}

	for _, ii := range testTable {
		t.Run(ii.testName, func(t *testing.T) {
			result, state := StringArrayToString(ii.array)

			if !state.IsOk() {
				t.Errorf(state.UserError)
				return
			}

			if result != ii.spected {
				t.Errorf("got %q, want %q", ii.spected, result)
			}

		})
	}
}
