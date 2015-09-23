package main

import "log"
import "sort"

type sortableStrings []string

func (str sortableStrings) Len() int {
	return len(str)
}

func (str sortableStrings) Less(i, j int) bool {
	return str[i] < str[j]
}

func (str sortableStrings) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
func main() {
	c := sortableStrings{"11", "22", "33", "00"}
	sort.Sort(c)
	log.Println(c)
}
