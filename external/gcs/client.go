package gcs

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Client struct {
	GcsClient *storage.Client
}

func InitGcsClient() *Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("external/gcs/practice-storage-v1-firebase-adminsdk-dwbs2-618816485c.json"))
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		GcsClient: client,
	}
}
