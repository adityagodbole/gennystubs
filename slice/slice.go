package slice

import "github.com/cheekybits/genny/generic"

type T generic.Type
type T1 generic.Type
type T2 generic.Type

type Slice_T []T
type Slice_T1 []T1
type Slice_T2 []T2

func (s Slice_T) Map(f func(e T) T) Slice_T {
	res := []T{}
	for _, e := range s {
		res = append(res, f(e))
	}
	return res
}

func (s Slice_T1) Map2(f func(e T1) T2) Slice_T2 {
	res := []T2{}
	for _, e := range s {
		res = append(res, f(e))
	}
	return res
}
