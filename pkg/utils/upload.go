package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"gohub/configs"
	"mime/multipart"
	"time"
)

func ImageUpload(fileHeader *multipart.FileHeader, folder string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	file, _ := fileHeader.Open()

	cfg := configs.GetConfig()
	cldService, _ := cloudinary.NewFromURL(cfg.UrlCloudinary)
	result, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folder})

	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
