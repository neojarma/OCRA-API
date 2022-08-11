package firebase_service

import (
	"context"
	"io"
)

type FirebaseService interface {
	CreateResource(ctx context.Context, path string, file io.Reader) error
	GenerateFirebaseStorageLink(path string) string
	DeleteResource(ctx context.Context, path string) error
}
