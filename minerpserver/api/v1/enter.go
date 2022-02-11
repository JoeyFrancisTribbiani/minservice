package v1

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/api/v1/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/api/v1/example"
	"github.com/JoeyFrancisTribbiani/minerpserver/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
