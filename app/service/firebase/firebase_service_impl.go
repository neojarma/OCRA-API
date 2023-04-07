package firebase_service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"sync"

	"cloud.google.com/go/storage"
)

const host string = "https://storage.googleapis.com/"

type FirebaseServiceImpl struct {
	Firestorage *storage.BucketHandle
}

func NewFirebaseService(firestorage *storage.BucketHandle) FirebaseService {
	var doOnce sync.Once
	service := new(FirebaseServiceImpl)

	doOnce.Do(func() {
		service = &FirebaseServiceImpl{
			Firestorage: firestorage,
		}
	})

	return service
}

func (service *FirebaseServiceImpl) CreateResource(ctx context.Context, path string, file *multipart.FileHeader) (string, error) {
	openFile, err := file.Open()
	if err != nil {
		return "", err
	}

	defer openFile.Close()

	path = service.getResourcePath(path, file.Filename)
	w := service.Firestorage.Object(path).NewWriter(ctx)
	if _, err := io.Copy(w, openFile); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	resouceLink := service.generateFirebaseStorageLink(path)
	return resouceLink, nil
}

func (service *FirebaseServiceImpl) getResourcePath(path, fileName string) string {
	newFileName := strings.ReplaceAll(fileName, " ", "")
	return path + newFileName
}

func (service *FirebaseServiceImpl) generateFirebaseStorageLink(path string) string {
	bucket := os.Getenv("FIREBASE_BUCKET_NAME")
	return fmt.Sprint(host, bucket, "/", path)
}

func (service *FirebaseServiceImpl) DeleteResource(ctx context.Context, path string) error {
	return service.Firestorage.Object(path).Delete(ctx)
}
