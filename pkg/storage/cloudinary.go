package storage

import (
	"context"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type StorageService interface {
	UploadFile(fileHeader *multipart.FileHeader, folder string) (string, error)
	DeleteFile(fileURL string) error
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

func (s *cloudinaryService) DeleteFile(fileURL string) error {
	publicID := extractPublicID(fileURL)
	if publicID == "" {
		return nil
	}

	_, err := s.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

func extractPublicID(fileURL string) string {
	idx := strings.Index(fileURL, "/upload/")
	if idx == -1 {
		return ""
	}
	path := fileURL[idx+len("/upload/"):]
	if strings.HasPrefix(path, "v") {
		if slashIdx := strings.Index(path, "/"); slashIdx != -1 {
			path = path[slashIdx+1:]
		}
	}
	if dotIdx := strings.LastIndex(path, "."); dotIdx != -1 {
		path = path[:dotIdx]
	}
	return path
}
