package uploads

import "mime/multipart"

type Repository interface {
	UploadLocal(file *multipart.FileHeader, nameFolder string) (string, error)
}