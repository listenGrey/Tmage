package models

type UploadImage struct {
	UserID     int64    `json:"user_id"`
	Tags       []string `json:"tags"`
	UploadTime string   `json:"upload_time"`
	ImageName  string   `json:"image_name"`
	Size       int64    `json:"size"`
	Type       string   `json:"type"`
	Content    []byte   `json:"content"`
}

type FormTags struct {
	Tags []string `json:"tags"`
}

type FormImageIds struct {
	ImageIds []string `json:"image_ids"`
}

type ModifyImage struct {
	UserID int64    `json:"user_id"`
	Tags   []string `json:"tags"`
}
