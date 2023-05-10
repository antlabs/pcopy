// Copyright [2020-2023] [guonaihong]
package dcopy

func Preheat(dst, src interface{}) error {
	return Copy(dst, src, withPreheat())
}
