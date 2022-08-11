package firebase_service

import (
	"context"
	"fmt"
	"io"
	"os"
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

func (service *FirebaseServiceImpl) CreateResource(ctx context.Context, path string, file io.Reader) error {
	w := service.Firestorage.Object(path).NewWriter(ctx)
	if _, err := io.Copy(w, file); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (service *FirebaseServiceImpl) GenerateFirebaseStorageLink(path string) string {
	bucket := os.Getenv("FIREBASE_BUCKET_NAME")
	return fmt.Sprint(host, bucket, "/", path)
}

func (service *FirebaseServiceImpl) DeleteResource(ctx context.Context, path string) error {
	return service.Firestorage.Object(path).Delete(ctx)
}
