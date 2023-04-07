package firebase_service

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func GetFirebaseStorageClient(ctx context.Context) *storage.BucketHandle {
	var doOnce sync.Once
	bucket := new(storage.BucketHandle)

	doOnce.Do(func() {
		option := getFirebaseConfiguration()
		client, err := storage.NewClient(ctx, option)
		if err != nil {
			log.Println(err.Error())
			return
		}

		bucketHandler := client.Bucket(os.Getenv("FIREBASE_BUCKET_NAME"))
		err = bucketHandler.DefaultObjectACL().Set(ctx, storage.AllUsers, storage.RoleReader)
		if err != nil {
			log.Println(err.Error())
			return
		}

		// CORS

		bucket = bucketHandler
	})

	return bucket
}

func getFirebaseConfiguration() option.ClientOption {
	return option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CONFIG")))
}
