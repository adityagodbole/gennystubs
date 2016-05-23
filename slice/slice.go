package slice

import "github.com/cheekybits/genny/generic"

type T generic.Type
type Slice_T []T

func (s Slice_T) Map(f func(e T) T) Slice_T {
	res := []T{}
	for _, e := range s {
		res = append(res, f(e))
	}
	return res
}
