package router

import (
	"minerpserver/router/autocode"
	"minerpserver/router/example"
	"minerpserver/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
