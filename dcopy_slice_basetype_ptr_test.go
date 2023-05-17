// Copyright [2020-2023] [guonaihong]
package dcopy

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type slicePtrTestSrc struct {
	A []*int
	B []*int8
	C []*int16
	D []*int32
	E []*int64
	F []*uint
	G []*uint8
	H []*uint16
	I []*uint32
	J []*uint64
	K []*uintptr
	L []*float32
	M []*float64
	N []*complex64
	O []*complex128
	P []*bool
	Q []*string
	// R []*interface{}
	S []*byte
	T []*rune
}

type slicePtrTestDst struct {
	A []int
	B []int8
	C []int16
	D []int32
	E []int64
	F []uint
	G []uint8
	H []uint16
	I []uint32
	J []uint64
	K []uintptr
	L []float32
	M []float64
	N []complex64
	O []complex128
	P []bool
	Q []string
	// R []interface{}
	S []byte
	T []rune
}

var (
	local_int_1 = 1
	local_int_2 = 2
	local_int_3 = 3
)

var (
	local_string_1 = "hello"
	local_string_2 = "world"
	local_string_3 = "!"
)

var (
	local_Complex64_1  = complex64(1)
	local_Complex64_2  = complex64(2)
	local_Complex64_3  = complex64(3)
	local_Complex128_1 = complex128(1)
	local_Complex128_2 = complex128(2)
	local_Complex128_3 = complex128(3)
)

var (
	local_bool_1 = true
	local_bool_2 = false
	local_bool_3 = true
)

var local_SlicePtrTestSrc = slicePtrTestSrc{
	A: []*int{&local_int_1, &local_int_2, &local_int_3},
	B: []*int8{(*int8)(unsafe.Pointer(&local_int_1)), (*int8)(unsafe.Pointer(&local_int_2)), (*int8)(unsafe.Pointer(&local_int_3))},
	C: []*int16{(*int16)(unsafe.Pointer(&local_int_1)), (*int16)(unsafe.Pointer(&local_int_2)), (*int16)(unsafe.Pointer(&local_int_3))},
	D: []*int32{(*int32)(unsafe.Pointer(&local_int_1)), (*int32)(unsafe.Pointer(&local_int_2)), (*int32)(unsafe.Pointer(&local_int_3))},
	E: []*int64{(*int64)(unsafe.Pointer(&local_int_1)), (*int64)(unsafe.Pointer(&local_int_2)), (*int64)(unsafe.Pointer(&local_int_3))},
	F: []*uint{(*uint)(unsafe.Pointer(&local_int_1)), (*uint)(unsafe.Pointer(&local_int_2)), (*uint)(unsafe.Pointer(&local_int_3))},
	G: []*uint8{(*uint8)(unsafe.Pointer(&local_int_1)), (*uint8)(unsafe.Pointer(&local_int_2)), (*uint8)(unsafe.Pointer(&local_int_3))},
	H: []*uint16{(*uint16)(unsafe.Pointer(&local_int_1)), (*uint16)(unsafe.Pointer(&local_int_2)), (*uint16)(unsafe.Pointer(&local_int_3))},
	I: []*uint32{(*uint32)(unsafe.Pointer(&local_int_1)), (*uint32)(unsafe.Pointer(&local_int_2)), (*uint32)(unsafe.Pointer(&local_int_3))},
	J: []*uint64{(*uint64)(unsafe.Pointer(&local_int_1)), (*uint64)(unsafe.Pointer(&local_int_2)), (*uint64)(unsafe.Pointer(&local_int_3))},
	K: []*uintptr{(*uintptr)(unsafe.Pointer(&local_int_1)), (*uintptr)(unsafe.Pointer(&local_int_2)), (*uintptr)(unsafe.Pointer(&local_int_3))},
	L: []*float32{(*float32)(unsafe.Pointer(&local_int_1)), (*float32)(unsafe.Pointer(&local_int_2)), (*float32)(unsafe.Pointer(&local_int_3))},
	M: []*float64{(*float64)(unsafe.Pointer(&local_int_1)), (*float64)(unsafe.Pointer(&local_int_2)), (*float64)(unsafe.Pointer(&local_int_3))},
	N: []*complex64{(*complex64)(unsafe.Pointer(&local_int_1)), (*complex64)(unsafe.Pointer(&local_int_2)), (*complex64)(unsafe.Pointer(&local_int_3))},
	O: []*complex128{(*complex128)(unsafe.Pointer(&local_int_1)), (*complex128)(unsafe.Pointer(&local_int_2)), (*complex128)(unsafe.Pointer(&local_int_3))},
	P: []*bool{(*bool)(unsafe.Pointer(&local_bool_1)), (*bool)(unsafe.Pointer(&local_bool_2)), (*bool)(unsafe.Pointer(&local_bool_3))},
	Q: []*string{(*string)(unsafe.Pointer(&local_string_1)), (*string)(unsafe.Pointer(&local_string_2)), (*string)(unsafe.Pointer(&local_string_3))},
	// R: []*interface{}{(*interface{})(unsafe.Pointer(&local_int_1)), (*interface{})(unsafe.Pointer(&local_int_2)), (*interface{})(unsafe.Pointer(&local_int_3))},
	S: []*byte{(*byte)(unsafe.Pointer(&local_int_1)), (*byte)(unsafe.Pointer(&local_int_2)), (*byte)(unsafe.Pointer(&local_int_3))},
	T: []*rune{(*rune)(unsafe.Pointer(&local_int_1)), (*rune)(unsafe.Pointer(&local_int_2)), (*rune)(unsafe.Pointer(&local_int_3))},
}

var local_SlicePtrTestDst = slicePtrTestDst{
	A: []int{local_int_1, local_int_2, local_int_3},
	B: []int8{int8(local_int_1), int8(local_int_2), int8(local_int_3)},
	C: []int16{int16(local_int_1), int16(local_int_2), int16(local_int_3)},
	D: []int32{int32(local_int_1), int32(local_int_2), int32(local_int_3)},
	E: []int64{int64(local_int_1), int64(local_int_2), int64(local_int_3)},
	F: []uint{uint(local_int_1), uint(local_int_2), uint(local_int_3)},
	G: []uint8{uint8(local_int_1), uint8(local_int_2), uint8(local_int_3)},
	H: []uint16{uint16(local_int_1), uint16(local_int_2), uint16(local_int_3)},
	I: []uint32{uint32(local_int_1), uint32(local_int_2), uint32(local_int_3)},
	J: []uint64{uint64(local_int_1), uint64(local_int_2), uint64(local_int_3)},
	K: []uintptr{uintptr(local_int_1), uintptr(local_int_2), uintptr(local_int_3)},
	L: []float32{float32(local_int_1), float32(local_int_2), float32(local_int_3)},
	M: []float64{float64(local_int_1), float64(local_int_2), float64(local_int_3)},
	N: []complex64{local_Complex64_1, local_Complex64_2, local_Complex64_3},
	O: []complex128{local_Complex128_1, local_Complex128_2, local_Complex128_3},
	P: []bool{local_int_1 != 0, local_int_2 != 0, local_int_3 != 0},
	Q: []string{local_string_1, local_string_2, local_string_3},
	// R: []interface{}{local_int_1, local_int_2, local_int_3},
	S: []byte{byte(local_int_1), byte(local_int_2), byte(local_int_3)},
	T: []rune{rune(local_int_1), rune(local_int_2), rune(local_int_3)},
}

func Test_Slice_BaseTypePtr_DstPtr_SrcPtr(t *testing.T) {
	err := Preheat(&slicePtrTestSrc{}, &slicePtrTestSrc{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}

		var dst slicePtrTestSrc
		err = Copy(&dst, &local_SlicePtrTestSrc, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %#v\n", dst.A)
		assert.Equal(t, dst.A, local_SlicePtrTestSrc.A)
		assert.Equal(t, dst.A[0], local_SlicePtrTestSrc.A[0])
		assert.Equal(t, dst.B, local_SlicePtrTestSrc.B)
		assert.Equal(t, dst.C, local_SlicePtrTestSrc.C)
		assert.Equal(t, dst.D, local_SlicePtrTestSrc.D)
		assert.Equal(t, dst.E, local_SlicePtrTestSrc.E)
		assert.Equal(t, dst.F, local_SlicePtrTestSrc.F)
		assert.Equal(t, dst.G, local_SlicePtrTestSrc.G)
		assert.Equal(t, dst.H, local_SlicePtrTestSrc.H)
		assert.Equal(t, dst.I, local_SlicePtrTestSrc.I)
		assert.Equal(t, dst.J, local_SlicePtrTestSrc.J)
		assert.Equal(t, dst.K, local_SlicePtrTestSrc.K)
		assert.Equal(t, dst.L, local_SlicePtrTestSrc.L)
		assert.Equal(t, dst.M, local_SlicePtrTestSrc.M)
		assert.Equal(t, dst.N, local_SlicePtrTestSrc.N)
		assert.Equal(t, dst.O, local_SlicePtrTestSrc.O)
		assert.Equal(t, dst.P, local_SlicePtrTestSrc.P)
		assert.Equal(t, dst.Q, local_SlicePtrTestSrc.Q)
		// assert.Equal(t, dst.R, local_SlicePtrTestSrc.R)
		assert.Equal(t, dst.S, local_SlicePtrTestSrc.S)
		assert.Equal(t, dst.T, local_SlicePtrTestSrc.T)

	}
}

func Test_Slice_BaseTypePtr_Dst_SrcPtr(t *testing.T) {
	err := Preheat(&slicePtrTestDst{}, &slicePtrTestSrc{})
	fmt.Printf("Preheat ok\n")
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}

		var dst slicePtrTestDst
		err = Copy(&dst, &local_SlicePtrTestSrc, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %#v\n", dst.A)
		for k := range dst.A {
			assert.Equal(t, dst.A[k], *local_SlicePtrTestSrc.A[k])
		}

		for k := range dst.B {
			assert.Equal(t, dst.B[k], *local_SlicePtrTestSrc.B[k])
		}

		for k := range dst.C {
			assert.Equal(t, dst.C[k], *local_SlicePtrTestSrc.C[k])
		}

		for k := range dst.D {
			assert.Equal(t, dst.D[k], *local_SlicePtrTestSrc.D[k])
		}

		for k := range dst.E {
			assert.Equal(t, dst.E[k], *local_SlicePtrTestSrc.E[k])
		}

		for k := range dst.F {
			assert.Equal(t, dst.F[k], *local_SlicePtrTestSrc.F[k])
		}

		for k := range dst.G {
			assert.Equal(t, dst.G[k], *local_SlicePtrTestSrc.G[k])
		}

		for k := range dst.H {
			assert.Equal(t, dst.H[k], *local_SlicePtrTestSrc.H[k])
		}

		for k := range dst.I {
			assert.Equal(t, dst.I[k], *local_SlicePtrTestSrc.I[k])
		}

		for k := range dst.J {
			assert.Equal(t, dst.J[k], *local_SlicePtrTestSrc.J[k])
		}

		for k := range dst.K {
			assert.Equal(t, dst.K[k], *local_SlicePtrTestSrc.K[k])
		}

		for k := range dst.L {
			assert.Equal(t, dst.L[k], *local_SlicePtrTestSrc.L[k])
		}

		for k := range dst.M {
			assert.Equal(t, dst.M[k], *local_SlicePtrTestSrc.M[k])
		}

		for k := range dst.N {
			assert.Equal(t, dst.N[k], *local_SlicePtrTestSrc.N[k])
		}

		for k := range dst.O {
			assert.Equal(t, dst.O[k], *local_SlicePtrTestSrc.O[k])
		}

		for k := range dst.P {
			assert.Equal(t, dst.P[k], *local_SlicePtrTestSrc.P[k])
		}

		for k := range dst.Q {
			assert.Equal(t, dst.Q[k], *local_SlicePtrTestSrc.Q[k])
		}

		// for k := range dst.R {
		// 	assert.Equal(t, dst.R[k], *local_SlicePtrTestSrc.R[k])
		// }

		for k := range dst.S {
			assert.Equal(t, dst.S[k], *local_SlicePtrTestSrc.S[k])
		}

		for k := range dst.T {
			assert.Equal(t, dst.T[k], *local_SlicePtrTestSrc.T[k])
		}

	}
}

func Test_Slice_BaseTypePtr_DstPtr_Src(t *testing.T) {
	err := Preheat(&slicePtrTestSrc{}, &slicePtrTestDst{})
	assert.NoError(t, err)

	var opts []Option
	for _, b := range []bool{true, false} {
		if b {
			opts = append(opts, WithUsePreheat())
		} else {
			opts = []Option{}
		}

		var dst slicePtrTestSrc
		err = Copy(&dst, &local_SlicePtrTestDst, opts...)
		assert.NoError(t, err)
		fmt.Printf("dst: %+v\n", dst)
		fmt.Printf(".A %#v\n", dst.A)
		for k := range local_SlicePtrTestDst.A {
			assert.Equal(t, *dst.A[k], local_SlicePtrTestDst.A[k])
		}

		for k := range dst.B {
			assert.Equal(t, *dst.B[k], local_SlicePtrTestDst.B[k])
		}

		for k := range dst.C {
			assert.Equal(t, *dst.C[k], local_SlicePtrTestDst.C[k])
		}

		for k := range dst.D {
			assert.Equal(t, *dst.D[k], local_SlicePtrTestDst.D[k])
		}

		for k := range dst.E {
			assert.Equal(t, *dst.E[k], local_SlicePtrTestDst.E[k])
		}

		for k := range dst.F {
			assert.Equal(t, *dst.F[k], local_SlicePtrTestDst.F[k])
		}

		for k := range dst.G {
			assert.Equal(t, *dst.G[k], local_SlicePtrTestDst.G[k])
		}

		for k := range dst.H {
			assert.Equal(t, *dst.H[k], local_SlicePtrTestDst.H[k])
		}

		for k := range dst.I {
			assert.Equal(t, *dst.I[k], local_SlicePtrTestDst.I[k])
		}

		for k := range dst.J {
			assert.Equal(t, *dst.J[k], local_SlicePtrTestDst.J[k])
		}

		for k := range dst.K {
			assert.Equal(t, *dst.K[k], local_SlicePtrTestDst.K[k])
		}

		for k := range dst.L {
			assert.Equal(t, *dst.L[k], local_SlicePtrTestDst.L[k])
		}

		for k := range dst.M {
			assert.Equal(t, *dst.M[k], local_SlicePtrTestDst.M[k])
		}

		for k := range dst.N {
			assert.Equal(t, *dst.N[k], local_SlicePtrTestDst.N[k])
		}

		for k := range dst.O {
			assert.Equal(t, *dst.O[k], local_SlicePtrTestDst.O[k])
		}

		for k := range dst.P {
			assert.Equal(t, *dst.P[k], local_SlicePtrTestDst.P[k])
		}

		for k := range dst.Q {
			assert.Equal(t, *dst.Q[k], local_SlicePtrTestDst.Q[k])
		}

		// for k := range dst.R {
		// 	assert.Equal(t, *dst.R[k], local_SlicePtrTestDst.R[k])
		// }

		for k := range dst.S {
			assert.Equal(t, *dst.S[k], local_SlicePtrTestDst.S[k])
		}

		for k := range dst.T {
			assert.Equal(t, *dst.T[k], local_SlicePtrTestDst.T[k])
		}
	}
}
