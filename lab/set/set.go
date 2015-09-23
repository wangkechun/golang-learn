package set

import (
	"bytes"
	"fmt"
)

// HashSet ...
type HashSet struct {
	m map[interface{}]bool
}

// New HashSet
func New() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// Add a element
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

// Remove a element
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

// Clear the HashSet
func (set *HashSet) Clear(e interface{}) {
	set.m = make(map[interface{}]bool)
}

//Contains a value
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//Len return length
func (set *HashSet) Len() int {
	return len(set.m)
}

//Same return true if other same now
func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// Elements return all keys
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

// IsSuperset return true if other in set
func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return true
	}
	otherLen := other.Len()
	setLen := other.Len()
	if setLen == 0 || otherLen == setLen {
		return false
	}
	if setLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// Set ...
type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}
