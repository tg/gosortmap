package sortmap

import (
	"reflect"
	"sort"
)

type KV struct {
	K, V interface{}
}

type Cmpr func(x, y KV) bool

type flatmap struct {
	kv []KV
	c  Cmpr
}

func newFlatMap(m interface{}, c Cmpr) *flatmap {
	fm := &flatmap{c: c}
	mv := reflect.ValueOf(m)
	for _, v := range mv.MapKeys() {
		fm.kv = append(fm.kv, KV{v.Interface(), mv.MapIndex(v).Interface()})
	}
	return fm
}

func (m flatmap) Len() int {
	return len(m.kv)
}
func (m flatmap) Less(i, j int) bool {
	return m.c(m.kv[i], m.kv[j])
}
func (m flatmap) Swap(i, j int) {
	m.kv[i], m.kv[j] = m.kv[j], m.kv[i]
}

func SortF(m interface{}, c Cmpr) []KV {
	fm := newFlatMap(m, c)
	sort.Sort(fm)
	return fm.kv
}

func SortValInt(m interface{}) []KV {
	return SortF(m, func(x, y KV) bool { return x.V.(int) < y.V.(int) })
}

func SortValIntDesc(m interface{}) []KV {
	return SortF(m, func(x, y KV) bool { return x.V.(int) > y.V.(int) })
}
