package response

import "github.com/JoeyFrancisTribbiani/minerpserver/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
