// Copyright [2020-2023] [guonaihong]
package pcopy

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var sliceElemIsPtrFuncTmpl = `
// Copyright [2020-2024] [guonaihong]
package pcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)
	// 自动生成的代码， 不要修改
	// 生成的代码位于，如下位置
  // set_basemap_tmpl.go
	// 生成命令位于 set_basemap_test.go
{{$TypeName := .TypeName}}
	{{range $_, $TypeVal := $TypeName}}
		{{$val := $TypeVal}}
		{{if eq  $val "int"}}
			{{$val = "reflect.Int"}}
		{{else if eq $val "int8"}}
			{{$val = "reflect.Int8"}}
		{{else if eq $val "int16"}}
			{{$val = "reflect.Int16"}}
		{{else if eq $val "int32"}}
			{{$val = "reflect.Int32"}}
		{{else if eq $val "int64"}}
			{{$val = "reflect.Int64"}}
		{{else if eq $val "uint"}}
			{{$val = "reflect.Uint"}}
		{{else if eq $val "uint8"}}
			{{$val = "reflect.Uint8"}}
		{{else if eq $val "uint16"}}
			{{$val = "reflect.Uint16"}}
		{{else if eq $val "uint32"}}
			{{$val = "reflect.Uint32"}}
		{{else if eq $val "uint64"}}
			{{$val = "reflect.Uint64"}}
		{{else if eq $val "float32"}}
			{{$val = "reflect.Float32"}}
		{{else if eq $val "float64"}}
			{{$val = "reflect.Float64"}}
		{{else if eq $val "string"}}
			{{$val = "reflect.String"}}
		{{else if eq $val "bool"}}
			{{$val = "reflect.Bool"}}
		{{else if eq $val "complex64"}}
			{{$val = "reflect.Complex64"}}
		{{else if eq $val "complex128"}}
			{{$val = "reflect.Complex128"}}
		{{else if eq $val "uintptr"}}
			{{$val = "reflect.Uintptr"}}
		{{end}}
		func setSliceElemIsBaseTypeOrPtr{{$TypeVal|title}}(dstType, srcType reflect.Type, dst, src unsafe.Pointer, opt options, of *offsetAndFunc) (err error) {
			if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == reflect.Ptr {
				srcSlice := *(*[]*{{$TypeVal}})(src)
				if len(srcSlice) == 0 {
					return nil
				}
				dstSlice := (*[]*{{$TypeVal}})(dst)
				if len(*dstSlice) < len(srcSlice) {
					*dstSlice = make([]*{{$TypeVal}}, 0, len(srcSlice))
				}
				for _, v := range srcSlice {
					var dv {{$TypeVal}}
					if v != nil {
						dv = *v
					}
					*dstSlice = append(*dstSlice, &dv)
				}
				return nil
			}

			if srcType.Elem().Kind() == {{$val}} && dstType.Elem().Kind() == reflect.Ptr {
				srcSlice := *(*[]{{$TypeVal}})(src)
				if len(srcSlice) == 0 {
					return nil
				}
				dstSlice := (*[]*{{$TypeVal}})(dst)
				if len(*dstSlice) < len(srcSlice) {
					*dstSlice = make([]*{{$TypeVal}}, 0, len(srcSlice))
				}
				for _, v := range srcSlice {
					dv := v
					*dstSlice = append(*dstSlice, &dv)
				}
				return nil
			}

			if srcType.Elem().Kind() == reflect.Ptr && dstType.Elem().Kind() == {{$val}} {
				srcSlice := *(*[]*{{$TypeVal}})(src)
				if len(srcSlice) == 0 {
					return nil
				}
				dstSlice := (*[]{{$TypeVal}})(dst)
				if len(*dstSlice) < len(srcSlice) {
					*dstSlice = make([]{{$TypeVal}}, 0, len(srcSlice))
				}
				for _, v := range srcSlice {
					var dv {{$TypeVal}}
					if v != nil {
						dv = *v
					}
					*dstSlice = append(*dstSlice, dv)
				}
				return nil
			}
			return nil
		}
{{end}}
`

var sliceElemIsPtrTable = `
{{$TypeName := .TypeName}}
var copySliceElemIsBaseTypeOrPtrTab = map[reflect.Kind]setReflectFunc{

{{range $_, $TypeVal := .TypeName}}
		{{$val := $TypeVal}}
		{{if eq  $val "int"}}
			{{$val = "reflect.Int"}}
		{{else if eq $val "int8"}}
			{{$val = "reflect.Int8"}}
		{{else if eq $val "int16"}}
			{{$val = "reflect.Int16"}}
		{{else if eq $val "int32"}}
			{{$val = "reflect.Int32"}}
		{{else if eq $val "int64"}}
			{{$val = "reflect.Int64"}}
		{{else if eq $val "uint"}}
			{{$val = "reflect.Uint"}}
		{{else if eq $val "uint8"}}
			{{$val = "reflect.Uint8"}}
		{{else if eq $val "uint16"}}
			{{$val = "reflect.Uint16"}}
		{{else if eq $val "uint32"}}
			{{$val = "reflect.Uint32"}}
		{{else if eq $val "uint64"}}
			{{$val = "reflect.Uint64"}}
		{{else if eq $val "float32"}}
			{{$val = "reflect.Float32"}}
		{{else if eq $val "float64"}}
			{{$val = "reflect.Float64"}}
		{{else if eq $val "string"}}
			{{$val = "reflect.String"}}
		{{else if eq $val "bool"}}
			{{$val = "reflect.Bool"}}
		{{else if eq $val "uintptr"}}
			{{$val = "reflect.Uintptr"}}
		{{else if eq $val "complex64"}}
			{{$val = "reflect.Complex64"}}
		{{else if eq $val "complex128"}}
			{{$val = "reflect.Complex128"}}
		{{end}}

	{{$val}}:     setSliceElemIsBaseTypeOrPtr{{$TypeVal|title}},
{{end}}
}

func getSetSliceElemIsBaseTypeOrPtrFunc(src reflect.Kind, p bool) setReflectFunc {
	f, ok := copySliceElemIsBaseTypeOrPtrTab[src]
	if p && !ok {
		panic(fmt.Sprintf("not support type:dst %v ", src))
	}
	return f
}
	`

// 基础类型表
var _sliceElemIsPtrTable = []string{
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
	"string",
	"bool",
	"uintptr",
	"complex64",
	"complex128",
}

func saveSliceElemIsPtrToFile(fileName string) error {
	var out bytes.Buffer
	tmpl, err := template.New("sliceElmeIsPtrFuncTmpl").Funcs(sprig.TxtFuncMap()).Parse(sliceElemIsPtrFuncTmpl)
	if err != nil {
		return fmt.Errorf("build sliceElemIsPtrFuncTmpl fail:%w", err)
	}

	err = tmpl.Execute(&out, baseTypeTmpl{TypeName: _sliceElemIsPtrTable})
	if err != nil {
		return err
	}
	tmpl, err = template.New("sliceElemIsPtrTable").Funcs(sprig.TxtFuncMap()).Parse(sliceElemIsPtrTable)
	if err != nil {
		return fmt.Errorf("build sliceElemIsPtrTableTmpl fail:%w", err)
	}
	err = tmpl.Execute(&out, baseTypeTmpl{TypeName: _sliceElemIsPtrTable})
	if err != nil {
		return err
	}

	// 格式化代码
	src, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("format code fail:%w", err)
	}

	return ioutil.WriteFile(fileName, src, 0o666)
}
