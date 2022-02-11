package router

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/router/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/router/example"
	"github.com/JoeyFrancisTribbiani/minerpserver/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
