package sqlStruct

import (
	"fmt"
	"go-project-example/tour/internal/word"
	"html/template"
	"os"
)

/**
 *@Author tudou
 *@Date 2020/7/22
 **/

const structTql = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`


type StructTemplate struct{
	structTql string
}


//存储Go结构体中所有的字段信息
type StructColumn struct{
	Name string
	Type string
	Tag string
	Comment string
}


//存储最终用于渲染的模板对象信息
type StructTemplateDB struct{
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate{
	return &StructTemplate{structTql: structTql}
}

//对查询COLUMNS得到的tbColumns进行分解和装华安
func (t *StructTemplate)AssemblyColumns(tbColumns []*TableColumn)[]*StructColumn{
	tplColumns := make([]*StructColumn,0,len(tbColumns))
	for _,column:=range tbColumns{
		tag:=fmt.Sprintf("`json:"+"\"%s\"`",column.ColumnName)
		tplColumns = append(tplColumns,&StructColumn{
			Name: column.ColumnName,
			Type: mySQLTypeToStructType[column.ColumnType],
			Tag: tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}


//处理完模板对象后，对模块渲染的自定义函数和模板对象进行处理
func (t *StructTemplate)Generate(tableName string,tplColumns []*StructColumn)(err error){
	tpl := template.Must(template.New("sqlStruct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTql))

	tplDB:=StructTemplateDB{
		TableName: tableName,
		Columns: tplColumns,
	}

	err = tpl.Execute(os.Stdout,tplDB)

	return
}