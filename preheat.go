// Copyright [2020-2023] [guonaihong]
package deepcopy

func Preheat(dst, src interface{}, opts ...Option) error {
	return CopyEx(dst, src, WithPreheat())
}
