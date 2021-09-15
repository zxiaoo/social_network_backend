package main

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

const (
	BUCKET_NAME = "xiao_around_bucket"
)

// read file from the request, save to GCS, and return a url for the file
func saveTOGCS(r io.Reader, objectName string) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}

	object := client.Bucket(BUCKET_NAME).Object(objectName)
	wc := object.NewWriter(ctx)
	// copy the file in r to the file on GCS
	if _, err := io.Copy(wc, r); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	// set access control for READ
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	// get attributes of the file saved on GCS
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}

	fmt.Printf("Image is saved to GCS: %s\n", attrs.MediaLink)
	return attrs.MediaLink, nil
}
