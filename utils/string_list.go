package utils

import (
	"slices"
	"sort"
	"strings"
)

// StringList is a list of string values
type StringList []string

// Empty indicates whether a change was found in this element
func (stringList *StringList) Empty() bool {
	return stringList == nil || len(*stringList) == 0
}

func (stringList *StringList) String() string {
	return strings.Join(*stringList, ", ")
}

func (stringList *StringList) Set(s string) error {
	*stringList = strings.Split(s, ",")
	return nil
}

func (stringList *StringList) Contains(s string) bool {
	if stringList == nil {
		return false
	}
	return slices.Contains(*stringList, s)
}

func (stringList *StringList) Minus(other StringList) StringList {
	return stringList.ToStringSet().Minus(other.ToStringSet()).ToStringList()
}

func (stringList *StringList) CartesianProduct(other StringList) []StringPair {
	result := make([]StringPair, stringList.Len()*other.Len())
	i := 0
	for _, a := range *stringList {
		for _, b := range other {
			result[i] = StringPair{a, b}
			i++
		}
	}
	return result
}

func (list StringList) ToStringSet() StringSet {
	result := make(StringSet, len(list))

	for _, s := range list {
		result[s] = struct{}{}
	}

	return result
}

func (list StringList) Sort() StringList {
	sort.Sort(list)
	return list
}

// Len implements the sort.Interface interface
func (list StringList) Len() int {
	return len(list)
}

// Less implements the sort.Interface interface
func (list StringList) Less(i, j int) bool {
	return list[i] < list[j]
}

// Swap implements the sort.Interface interface
func (list StringList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (stringList *StringList) Is(s string) bool {
	if stringList == nil {
		return false
	}
	return len(*stringList) == 1 && (*stringList)[0] == s
}
