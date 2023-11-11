package datafileprocess

import domainerror "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/error"

type DatafileUploadEvent struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type DatafileProcessServiceInterface interface {
	Process(DatafileUploadEvent) *domainerror.DomainError
}
