package setHelper

/*
	This implementation of sets was taken from:
	https://github.com/golang-collections/collections/tree/604e922904d35e97f98a774db7881f049cd8d970/set
	Documentation:
	https://pkg.go.dev/github.com/golang-collections/collections/set
*/

// Set is an implementation of set
type (
	Set struct {
		hash map[interface{}]nothing
	}

	nothing struct{}
)

// New fucntion creates a new set
func New(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]nothing)}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

// Do method calls foo function for each item in the set
func (set *Set) Do(foo func(interface{})) {
	for k := range set.hash {
		foo(k)
	}
}

// Has method verifies if an element is or not the element in the set
func (set *Set) Has(element interface{}) bool {
	_, exists := set.hash[element]
	return exists
}

// Insert method insert an element to the set
func (set *Set) Insert(element interface{}) {
	set.hash[element] = nothing{}
}

// Length method return the number of items in the set
func (set *Set) Length() int {
	return len(set.hash)
}

// Remove method remove an element from the set
func (set *Set) Remove(element interface{}) {
	delete(set.hash, element)
}
