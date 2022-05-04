package utils_test

import (
	"context"
	"testing"

	"github.com/KapitanD/cyan-project/internal/constants"
	"github.com/KapitanD/cyan-project/pkg/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func Test_UnaryValueFromMetadata(t *testing.T) {
	testCases := []struct {
		Name      string
		Context   context.Context
		Key       constants.MdKey
		Value     string
		ExpectErr error
	}{
		{
			Name: "ok with token",
			Context: metadata.NewIncomingContext(
				context.Background(),
				metadata.Pairs(string(constants.MdKeyAuthToken), "sometoken"),
			),
			Key:       constants.MdKeyAuthToken,
			Value:     "sometoken",
			ExpectErr: nil,
		},
		{
			Name:      "context with no metadata",
			Context:   context.Background(),
			Key:       constants.MdKeyAuthToken,
			Value:     "",
			ExpectErr: utils.EmptyMetaErr,
		},
		{
			Name: "context with meta but wrong key",
			Context: metadata.NewIncomingContext(
				context.Background(),
				metadata.Pairs("boken", "sometoken"),
			),
			Key:       constants.MdKeyAuthToken,
			Value:     "",
			ExpectErr: utils.InvalidValLenErr,
		},
		{
			Name: "context with 2 values for needed key",
			Context: metadata.NewIncomingContext(
				context.Background(),
				metadata.Pairs(
					string(constants.MdKeyAuthToken), "value1",
					string(constants.MdKeyAuthToken), "value2",
				),
			),
			Key:       constants.MdKeyAuthToken,
			Value:     "",
			ExpectErr: utils.InvalidValLenErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			val, err := utils.UnaryValueFromMetadata(tc.Context, tc.Key)
			assert.ErrorIs(t, err, tc.ExpectErr)

			assert.Equal(t, tc.Value, val)
		})
	}
}
