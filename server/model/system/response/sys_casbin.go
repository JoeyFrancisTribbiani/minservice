package response

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
