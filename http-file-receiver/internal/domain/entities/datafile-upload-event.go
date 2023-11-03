package entities

type DatafileUploadEvent struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}
