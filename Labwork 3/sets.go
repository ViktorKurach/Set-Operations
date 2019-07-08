package sets

import "strconv"

type Set []string

var EmptySet = Set{}

func maybeAppend(el string, set Set) Set {
	for _, x := range set {
		if x == el {
			return set
		}
	}
	return append(set, el)
}

func isInSet(el string, set Set) bool {
	for _, x := range set {
		if len(el) == len(x) && x == el {
			return true
		}
	}
	return false
}

func createSet(elems []string) Set {
	res := EmptySet
	for _, el := range elems {
		res = maybeAppend(el, res)
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

// Returns a set serialized into string
func SerializeSet(set Set) string {
	res := "["
	for _, el := range set {
		res += el + ", "
	}
	return res[0:len(res)-2] + "]"
}

// Returns true if set1 is subset of set2, or false otherwise
func IsSubset(set1 Set, set2 Set) bool {
	for _, el := range set1 {
		if !isInSet(el, set2) {
			return false
		}
	}
	return true
}

// Returns true if both sets have same elements, or false otherwise
func EqualSets(set1 Set, set2 Set) bool {
	return IsSubset(set1, set2) && IsSubset(set2, set1)
}

// Intersection of sets: elements appearing both in set1 and set2
func Intersection(set1 Set, set2 Set) Set {
	res := EmptySet
	for _, el1 := range set1 {
		for _, el2 := range set2 {
			if el1 == el2 {
				res = append(res, el1)
				break
			}
		}
	}
	return res
}

// Union of sets: elements appearing either in set1 or set2
func Union(set1 Set, set2 Set) Set {
	res := set1
	for _, el := range set2 {
		res = maybeAppend(el, res)
	}
	return res
}

// Difference of sets: elements appearing in set1, but not in set2
func Difference(set1 Set, set2 Set) Set {
	res := EmptySet
	var mayAppend = true
	for _, el1 := range set1 {
		mayAppend = true
		for _, el2 := range set2 {
			if el1 == el2 {
				mayAppend = false
				break
			}
		}
		if mayAppend {
			res = append(res, el1)
		}
	}
	return res
}
