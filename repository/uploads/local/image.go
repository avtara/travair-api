package local

import (
	"github.com/avtara/travair-api/businesses/uploads"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type uploadRepository struct {
	rootFolder string
	baseUrl    string
}

func NewUploadRepository(rf string, bu string) uploads.Repository {
	return &uploadRepository{
		rootFolder: rf,
		baseUrl:    bu,
	}
}

func (us *uploadRepository) UploadLocal(file *multipart.FileHeader, destFolder string) (string, error) {
	sourceFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	folderPath := us.rootFolder + "/" + destFolder
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	destinationFile := folderPath + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + file.Filename
	destination, err := os.Create(destinationFile)
	if err != nil {
		return "", err
	}

	defer destination.Close()

	if _, err = io.Copy(destination, sourceFile); err != nil {
		return "", err
	}

	return us.baseUrl + "/" + destinationFile, nil
}
