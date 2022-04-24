package utils

import (
	"context"

	"github.com/KapitanD/cyan-project/internal/constants"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

func UnaryValueFromMetadata(ctx context.Context, key constants.MdKey) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("empty metadata")
	}

	mdValue := md.Get(string(key))
	if len(mdValue) != 1 {
		return "", errors.New("invalid len of email value")
	}

	return mdValue[0], nil
}
