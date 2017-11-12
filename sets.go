package sets

import "strconv"

type Set []string

var EmptySet = Set{}

func isInSet(el string, set Set) bool {
	for _, x := range set {
		if x == el {
			return true
		}
	}
	return false
}

func createSet(elems []string) Set {
	res := EmptySet
	for _, el := range elems {
		if !isInSet(el, res) {
			res = append(res, el)
		}
	}
	return res
}

func convertInt(elem int64) string {
	return strconv.FormatInt(elem, 10)
}

func convertUint(elem uint64) string {
	return strconv.FormatUint(elem, 10)
}

func convertFloat(elem float64) string {
	return strconv.FormatFloat(elem, 'f', 10, 64)
}

func convertBool(elem bool) string {
	return strconv.FormatBool(elem)
}

// Returns true if integer belongs to set, or false otherwise
func IsIntInSet(elem int64, set Set) bool {
	return isInSet(convertInt(elem), set)
}

// Returns true if unsigned integer belongs to set, or false otherwise
func IsUintInSet(elem uint64, set Set) bool {
	return isInSet(convertUint(elem), set)
}

// Returns true if float belongs to set, or false otherwise
func IsFloatInSet(elem float64, set Set) bool {
	return isInSet(convertFloat(elem), set)
}

// Returns true if bool belongs to set, or false otherwise
func IsBoolInSet(elem bool, set Set) bool {
	return isInSet(convertBool(elem), set)
}

// Returns true if string belongs to set, or false otherwise
func IsStringInSet(elem string, set Set) bool {
	return isInSet(elem, set)
}

// Creates set of integers
func CreateSetOfInts(elems []int64) Set {
	converted := []string{}
	for _, el := range elems {
		converted = append(converted, convertInt(el))
	}
	return createSet(converted)
}

// Creates set of unsigned integers
func CreateSetOfUints(elems []uint64) Set {
	converted := []string{}
	for _, el := range elems {
		converted = append(converted, convertUint(el))
	}
	return createSet(converted)
}

// Creates set of floats, truncating them to 10 digits after comma
func CreateSetOfFloats(elems []float64) Set {
	converted := []string{}
	for _, el := range elems {
		converted = append(converted, convertFloat(el))
	}
	return createSet(converted)
}

// Creates set of bools
func CreateSetOfBools(elems []bool) Set {
	converted := []string{}
	for _, el := range elems {
		converted = append(converted, convertBool(el))
	}
	return createSet(converted)
}

// Creates set of strings
func CreateSetOfStrings(elems []string) Set {
	return createSet(elems)
}

// Returns true if both sets have same elements, or false otherwise
func EqualSets(set1 Set, set2 Set) bool {
	for _, el := range set1 {
		if !isInSet(el, set2) {
			return false
		}
	}
	return true
}

// Intersection of sets: elements appearing both in set1 and set2
func Intersection(set1 Set, set2 Set) Set {
	res := EmptySet
	for _, el := range set1 {
		if isInSet(el, set2) {
			res = append(res, el)
		}
	}
	return res
}

// Union of sets: elements appearing either in set1 or set2
func Union(set1 Set, set2 Set) Set {
	res := set1
	for _, el := range set2 {
		if !isInSet(el, set1) {
			res = append(res, el)
		}
	}
	return res
}

// Difference of sets: elements appearing in set1, but not in set2
func Difference(set1 Set, set2 Set) Set {
	res := EmptySet
	for _, el := range set1 {
		if !isInSet(el, set2) {
			res = append(res, el)
		}
	}
	return res
}
