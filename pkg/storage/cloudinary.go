package storage

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type StorageService interface {
	UploadFile(fileHeader *multipart.FileHeader, folder string) (string, error)
}

type cloudinaryService struct{
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() (StorageService, error){
	cldURL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cldURL)
	if err != nil {
		return nil, err
	}
	return &cloudinaryService{cld: cld}, nil
}

func (s *cloudinaryService) UploadFile(fileHeader *multipart.FileHeader, folder string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	uploadRes, err := s.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", err
	}
	return uploadRes.SecureURL, nil
}