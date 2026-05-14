package core_minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (c *Client) DeleteObject(ctx context.Context, objectName string) error {
	err := c.cli.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("remove object: %w", err)
	}

	return nil
}
