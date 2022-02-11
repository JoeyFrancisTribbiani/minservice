package response

import "minerpserver/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
