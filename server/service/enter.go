package service

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/service/autocode"
	"github.com/JoeyFrancisTribbiani/minerpserver/service/example"
	"github.com/JoeyFrancisTribbiani/minerpserver/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
