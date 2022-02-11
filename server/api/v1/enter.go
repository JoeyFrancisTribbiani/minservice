package v1

import (
	"minerpserver/api/v1/autocode"
	"minerpserver/api/v1/example"
	"minerpserver/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
