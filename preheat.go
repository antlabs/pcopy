// Copyright [2020-2023] [guonaihong]
package pcopy

func Preheat[T any, U any](dst *T, src *U) error {
	return Copy(dst, src, withPreheat())
}
