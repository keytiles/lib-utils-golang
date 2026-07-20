package kt_utils

import (
	"github.com/keytiles/lib-sets-golang/v2/pkg/kt_sets"
)

// `maps` package returns an iterable seq only of keys of a map - this one returns a real slice.
// Of course this also works: `keys := slices.Collect(maps.Keys(m))` but this one might be more efficient as it
// is allocating slice based on size of the map immediately - the above can not.
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	return keys
}

// `maps` package returns an iterable seq only of values of a map - this one returns values as a Set.
// Of course you can also do this `values := slices.Collect(maps.Values(m))` but that returns a slice and duplicated
// values appear there as is.
func MapValuesSet[K comparable, V comparable](m map[K]V) *kt_sets.Set[V] {
	set := kt_sets.NewSetWithCapacity[V](len(m))
	for _, v := range m {
		set.Add(v)
	}
	return set
}
