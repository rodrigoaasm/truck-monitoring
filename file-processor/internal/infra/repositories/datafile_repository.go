package repositories

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

func (repo MinIORepository) GetDatafile(filename string) (io.Reader, error) {
	return repo.minioClient.GetObject(repo.context, "truck-files", filename, minio.GetObjectOptions{})
}
