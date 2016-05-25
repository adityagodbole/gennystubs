package slice

import "github.com/cheekybits/genny/generic"

type T generic.Type
type Slice_T []T

func (sl Slice_T) Map(f func(e T) T) Slice_T {
	return Map_T(sl, f)
}

func (sl Slice_T) Filter(f func(e T) bool) Slice_T {
	return Filter_T(sl, f)
}

func (sl Slice_T) Contains(o T) (bool, int) {
	return Contains_T(sl, o)
}

func (sl Slice_T) Each(f func(e T)) {
	for _, e := range sl {
		f(e)
	}
}

func Filter_T(sl []T, f func(e T) bool) []T {
	res := []T{}
	for _, e := range sl {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

func Map_T(sl []T, f func(e T) T) []T {
	res := []T{}
	for _, e := range sl {
		res = append(res, f(e))
	}
	return res
}

func Contains_T(sl []T, o T) (bool, int) {
	for i, v := range sl {
		if v == o {
			return true, i
		}
	}
	return false, -1
}
