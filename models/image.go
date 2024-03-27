package models

import "mime/multipart"

type UploadImage struct {
	UserID     int64
	Tags       []string
	UploadTime string
	Image      *multipart.FileHeader
}
