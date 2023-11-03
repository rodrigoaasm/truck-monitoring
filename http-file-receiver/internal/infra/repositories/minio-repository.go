package repositories

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/entities"
)

type MinIORepository struct {
	minioClient *minio.Client
	context     context.Context
}

func NewMinIORepository(endpoint, accessKeyID, secretAccessKey string) *MinIORepository {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &MinIORepository{
		minioClient: minioClient,
		context:     context.Background(),
	}
}

func (repo MinIORepository) PutDatafile(file io.Reader, datafileUploadEvent entities.DatafileUploadEvent) error {
	_, err := repo.minioClient.PutObject(repo.context, "truck-files", datafileUploadEvent.Filename, file, datafileUploadEvent.Size, minio.PutObjectOptions{})

	return err
}
