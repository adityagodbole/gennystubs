package slice

import "github.com/cheekybits/genny/generic"

type T generic.Type
type T1 generic.Type

type Slice_T []T
type Slice_T1 []T1

func (s Slice_T) Map(f func(e T) T) Slice_T {
	res := []T{}
	for _, e := range s {
		res = append(res, f(e))
	}
	return res
}

func (s Slice_T) Map_T1(f func(e T) T1) Slice_T1 {
	res := []T1{}
	for _, e := range s {
		res = append(res, f(e))
	}
	return res
}
