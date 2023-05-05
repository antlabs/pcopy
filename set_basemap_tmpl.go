package deepcopy

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"

	"github.com/Masterminds/sprig/v3"
)

var baseMapFuncTmpl = `
package deepcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)
	// 自动生成的代码， 不要修改
	// 生成命令位于 set_basemap_test.go
{{$TypeName := .TypeName}}
{{range $_, $TypeKey := .TypeName}}
	{{range $_, $TypeVal := $TypeName}}
		func setBaseMap{{$TypeKey|title}}{{$TypeVal|title}}(dstAddr, srcAddr unsafe.Pointer) {
			src := (*map[{{$TypeKey}}]{{$TypeVal}})(srcAddr)
			if len(*src) == 0 {
				return	
			}
			dst := (*map[{{$TypeKey}}]{{$TypeVal}})(dstAddr)
			if dst == nil || len(*dst) == 0 {
				*dst = make(map[{{$TypeKey}}]{{$TypeVal}}, len(*src))
			}

			for k, v := range *src {
				(*dst)[k] = v
			}
		}
	{{end}}
{{end}}
`

var baseMapTable = `
{{$TypeName := .TypeName}}
var baseMapTab = setBaseMapFuncTab{

{{range $_, $TypeKey := .TypeName}}
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
		{{end}}

		{{$key := $TypeKey}}
		{{if eq  $key "int"}}
			{{$key = "reflect.Int"}}
		{{else if eq $key "int8"}}
			{{$key = "reflect.Int8"}}
		{{else if eq $key "int16"}}
			{{$key = "reflect.Int16"}}
		{{else if eq $key "int32"}}
			{{$key = "reflect.Int32"}}
		{{else if eq $key "int64"}}
			{{$key = "reflect.Int64"}}
		{{else if eq $key "uint"}}
			{{$key = "reflect.Uint"}}
		{{else if eq $key "uint8"}}
			{{$key = "reflect.Uint8"}}
		{{else if eq $key "uint16"}}
			{{$key = "reflect.Uint16"}}
		{{else if eq $key "uint32"}}
			{{$key = "reflect.Uint32"}}
		{{else if eq $key "uint64"}}
			{{$key = "reflect.Uint64"}}
		{{else if eq $key "float32"}}
			{{$key = "reflect.Float32"}}
		{{else if eq $key "float64"}}
			{{$key = "reflect.Float64"}}
		{{else if eq $key "string"}}
			{{$key = "reflect.String"}}
		{{else if eq $key "bool"}}
			{{$key = "reflect.Bool"}}
		{{end}}

	baseMapKind{key:{{$key}}, val:{{$val}}}:       setBaseMap{{$TypeKey|title}}{{$TypeVal|title}},
	{{end}}
{{end}}
}

func getSetBaseMapFunc(key reflect.Kind, val reflect.Kind) setUnsafeFunc {
	k := baseMapKind{key: key, val: val}
	f, ok := baseMapTab[k]
	if !ok {
		panic(fmt.Sprintf("not support type:key %v val: %v", key, val))
	}
	return f
}
	`

// 基础类型表
var baseTypeTable = []string{
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
}

func saveBaseMapFuncToFile(fileName string) error {
	var out bytes.Buffer
	tmpl, err := template.New("setBaseMapFunc").Funcs(sprig.FuncMap()).Parse(baseMapFuncTmpl)
	if err != nil {
		return fmt.Errorf("build setBaseMapFunc fail:%w", err)
	}

	err = tmpl.Execute(&out, baseMapTypeTmpl{TypeName: baseTypeTable})
	if err != nil {
		return err
	}
	tmpl, err = template.New("setBaseMapTable").Funcs(sprig.FuncMap()).Parse(baseMapTable)
	if err != nil {
		return fmt.Errorf("build setBaseMapTable fail:%w", err)
	}
	err = tmpl.Execute(&out, baseMapTypeTmpl{TypeName: baseTypeTable})
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
