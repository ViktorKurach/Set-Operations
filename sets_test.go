package sets

import (
	"math"
	"testing"
)

type setAndSet struct {
	set1     Set
	set2     Set
	interRes Set
	unionRes Set
	difRes   Set
}

var testCases = []setAndSet{
	{
		CreateSetOfInts([]int64{1, 2, 2, 2}),
		CreateSetOfInts([]int64{-1, 1}),
		CreateSetOfInts([]int64{1}),
		CreateSetOfInts([]int64{1, 2, -1}),
		CreateSetOfInts([]int64{2}),
	},
	{
		CreateSetOfUints([]uint64{1, 2}),
		CreateSetOfUints([]uint64{3, 4}),
		EmptySet,
		CreateSetOfUints([]uint64{1, 2, 3, 4}),
		CreateSetOfUints([]uint64{1, 2}),
	},
	{
		CreateSetOfFloats([]float64{math.Pi, math.Phi, math.E}),
		CreateSetOfFloats([]float64{math.Pi}),
		CreateSetOfFloats([]float64{math.Pi}),
		CreateSetOfFloats([]float64{math.Pi, math.Phi, math.E}),
		CreateSetOfFloats([]float64{math.Phi, math.E}),
	},
	{
		CreateSetOfStrings([]string{"bar", "baz"}),
		CreateSetOfStrings([]string{"foo", "bar", "baz"}),
		CreateSetOfStrings([]string{"bar", "baz"}),
		CreateSetOfStrings([]string{"foo", "bar", "baz"}),
		EmptySet,
	},
	{
		CreateSetOfBools([]bool{true, !false, true || false}),
		CreateSetOfBools([]bool{true && false}),
		EmptySet,
		CreateSetOfBools([]bool{true, false}),
		CreateSetOfBools([]bool{true}),
	},
	{
		CreateSetOfUints([]uint64{0, 1, 5}),
		CreateSetOfStrings([]string{"1", "2"}),
		CreateSetOfStrings([]string{"1"}),
		CreateSetOfStrings([]string{"0", "5", "1", "2"}),
		CreateSetOfStrings([]string{"0", "5"}),
	},
	{
		CreateSetOfBools([]bool{true, true && true}),
		EmptySet,
		EmptySet,
		CreateSetOfBools([]bool{true}),
		CreateSetOfBools([]bool{true}),
	},
	{
		EmptySet,
		EmptySet,
		EmptySet,
		EmptySet,
		EmptySet,
	},
}

func extendTestCases(testCases []setAndSet, N uint64) []setAndSet {
	testCasesLen := uint64(len(testCases))
	for i := uint64(0); i < N; i++ {
		for j := uint64(0); j < testCasesLen; j++ {
			testCases = append(testCases, testCases[j])
		}
	}
	return testCases
}

func TestSetOperations(t *testing.T) {
	extTextCases := extendTestCases(testCases, uint64(1000000))
	for _, tCase := range extTextCases {
		if !EqualSets(Intersection(tCase.set1, tCase.set2), tCase.interRes) {
			t.Error(tCase.set1, "ᴖ", tCase.set2, "!=", tCase.interRes)
		}
		if !EqualSets(Union(tCase.set1, tCase.set2), tCase.unionRes) {
			t.Error(tCase.set1, "ᴗ", tCase.set2, "!=", tCase.unionRes)
		}
		if !EqualSets(Difference(tCase.set1, tCase.set2), tCase.difRes) {
			t.Error(tCase.set1, "\\", tCase.set2, "!=", tCase.difRes)
		}
	}
}
