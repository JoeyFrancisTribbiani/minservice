package response

import "github.com/JoeyFrancisTribbiani/minerpserver/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
