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

type ShareImage struct {
	UserID         int64    `json:"user_id"`
	ImagesIds      []string `json:"images_ids"`
	Token          string   `json:"encoded_token"`
	ShareTime      string   `json:"share_time"`
	ExpirationTime string   `json:"expiration_time"`
}
