package dao

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"
)

func GetSetInt64List(ctx context.Context, key string) ([]uint64, error) {
	res, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, len(res))
	for i, v := range res {
		ids[i] = convert.StrTo(v).MustUInt64()
	}

	return ids, nil
}
