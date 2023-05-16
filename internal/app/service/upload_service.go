package service

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadService struct {
	CloudName string
	APIKey    string
	APISecret string
	Directory string
}

func NewUploadService(CloudName, APIKey, APISecret, Directory string) *UploadService {
	return &UploadService{
		CloudName: CloudName,
		APIKey:    APIKey,
		APISecret: APISecret,
		Directory: Directory,
	}
}

func (us *UploadService) UploadImage(input *multipart.FileHeader) (string, error) {
	cld, _ := cloudinary.NewFromParams(us.CloudName, us.APIKey, us.APISecret)

	ctx := context.Background()
	file, _ := input.Open()

	uploaded, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			Folder: us.CloudName,
		},
	)
	if err != nil {
		return "", err
	}

	return uploaded.SecureURL, nil
}
