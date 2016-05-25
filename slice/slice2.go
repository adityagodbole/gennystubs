package slice

import "github.com/cheekybits/genny/generic"

type T generic.Type
type V generic.Type

func Map_T_V(sl []T, f func(e T) V) []V {
	res := []V{}
	for _, e := range sl {
		res = append(res, f(e))
	}
	return res
}
