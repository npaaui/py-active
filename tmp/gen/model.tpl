package {{.PackageName}}
{{$exportModelName := .ModelName | FormatCamelcase}}
import (
    . "active/common"
)

/**{{range .TableSchema}}
"{{.Field}}": "{{.Type | TypeConvert}}", // {{.Comment}} {{end}}
 */

type {{$exportModelName}} struct {
{{range .TableSchema}}    {{.Field | ExportColumn | FormatCamelcase}} {{.Type | TypeConvert}} {{.Field | Tags}}
{{end}}}

func New{{$exportModelName}}Model() *{{$exportModelName}} {
	return &{{$exportModelName}}{}
}

func (m *{{$exportModelName}}) Info() bool {
	has, err := GetDbEngineIns().Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *{{$exportModelName}}) Insert() int64 {
	row, err := GetDbEngineIns().Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *{{$exportModelName}}) Update(arg *{{$exportModelName}}) int64 {
	row, err := GetDbEngineIns().Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *{{$exportModelName}}) Delete() int64 {
	row, err := GetDbEngineIns().Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

{{range .TableSchema}}
func (m *{{$exportModelName}}) Set{{.Field | FormatCamelcase}}(arg {{.Type | TypeConvert}}) *{{$exportModelName}} {
	m.{{.Field | FormatCamelcase}} = arg
	return m
}
{{end}}
func (m {{$exportModelName}}) AsMapItf() MapItf {
	return MapItf{ {{range .TableSchema}}
        "{{.Field}}": m.{{.Field | FormatCamelcase}}, {{end}}
	}
}
func (m {{$exportModelName}}) Translates() map[string]string {
	return map[string]string{ {{range .TableSchema}}
        "{{.Field}}": "{{.Comment}}", {{end}}
	}
}