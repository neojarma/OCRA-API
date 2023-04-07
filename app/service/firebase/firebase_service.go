package firebase_service

import (
	"context"
	"mime/multipart"
)

type FirebaseService interface {
	CreateResource(ctx context.Context, path string, file *multipart.FileHeader) (string, error)
	generateFirebaseStorageLink(path string) string
	DeleteResource(ctx context.Context, path string) error
}
