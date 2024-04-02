package models

type UploadImage struct {
	UserID     int64    `json:"user_id"`
	Tags       []string `json:"tags"`
	UploadTime string   `json:"upload_time"`
	ImageName  string   `json:"image_name"`
	Size       int64    `json:"size"`
	Content    []byte   `json:"content"`
}
