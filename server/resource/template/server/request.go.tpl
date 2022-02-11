package request

import (
	"minerpserver/model/autocode"
	"minerpserver/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}