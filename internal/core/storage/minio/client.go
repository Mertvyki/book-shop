package core_minio

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	cli           *minio.Client
	bucket        string
	publicBaseURL string
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	cli, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("create minio client: %w", err)
	}

	exists, err := cli.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("check bucket exists: %w", err)
	}
	if !exists {
		if err := cli.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("create bucket: %w", err)
		}
	}

	return &Client{
		cli:           cli,
		bucket:        cfg.Bucket,
		publicBaseURL: strings.TrimRight(cfg.PublicBaseURL, "/"),
	}, nil
}

func (c *Client) Upload(
	ctx context.Context,
	objectName string,
	r io.Reader,
	size int64,
	contentType string,
) (string, error) {
	_, err := c.cli.PutObject(ctx, c.bucket, objectName, r, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("put object: %w", err)
	}

	return objectName, nil
}

func (c *Client) PresignedGetObject(
	ctx context.Context,
	objectName string,
	expiry time.Duration,
) (string, error) {

	url, err := c.cli.PresignedGetObject(
		ctx,
		c.bucket,
		objectName,
		expiry,
		nil,
	)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (c *Client) GetObject(
	ctx context.Context,
	objectName string,
) (io.ReadCloser, string, error) {
	obj, err := c.cli.GetObject(ctx, c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", fmt.Errorf("get object: %w", err)
	}

	stat, err := obj.Stat()
	if err != nil {
		obj.Close()
		return nil, "", fmt.Errorf("stat object: %w", err)
	}

	return obj, stat.ContentType, nil
}
