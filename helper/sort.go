package helper

import (
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func mapToPairList(mymap map[string]int) PairList {
	pl := make(PairList, len(mymap))
	i := 0
	for k, v := range mymap {
		pl[i] = Pair{k, v}
		i++
	}
	return pl
}

func pairListToSlice(pl PairList) []string {
	result := []string{}

	for _, p := range pl {
		result = append(result, p.Key)
	}
	return result
}

func KeysSortedByValues(mymap map[string]int) []string {
	pl := mapToPairList(mymap)
	sort.Sort(pl)
	return pairListToSlice(pl)
}

func KeysSortedByValuesReverse(mymap map[string]int) []string {
	pl := mapToPairList(mymap)

	sort.Sort(sort.Reverse(pl))

	return pairListToSlice(pl)
}

func KeyWithBiggestValue(mymap map[string]int) string {
	key := ""
	value := 0

	for k, v := range mymap {
		if v > value {
			value = v
			key = k
		}
	}

	return key
}

func GetBiggestValueInMap(mymap map[string]int) int {

	value := 0

	for _, v := range mymap {
		if v > value {
			value = v

		}
	}

	return value
}

func ReverseStringSlice(src []string) []string {
	result := []string{}

	for i := len(src) - 1; i >= 0; i-- {
		result = append(result, src[i])
	}

	return result
}
