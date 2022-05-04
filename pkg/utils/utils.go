package utils

import (
	"context"

	"github.com/KapitanD/cyan-project/internal/constants"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

var (
	EmptyMetaErr     = errors.New("empty metadata")
	InvalidValLenErr = errors.New("invalid len of key value")
)

func UnaryValueFromMetadata(ctx context.Context, key constants.MdKey) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", EmptyMetaErr
	}

	mdValue := md.Get(string(key))
	if len(mdValue) != 1 {
		return "", InvalidValLenErr
	}

	return mdValue[0], nil
}
