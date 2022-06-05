package dao

import (
	"context"
	"fmt"

	"github.com/weirdo0314/tiny-tiktok/pkg/constants"

	"go.uber.org/zap"
)

func QueryFavoriteCount(ctx context.Context, vid uint64) uint64 {
	count, err := rdb.SCard(ctx, fmt.Sprintf("favorite:%d", vid)).Result()
	if err != nil {
		zap.L().Error("redis query follower count error", zap.Uint64("video id", vid), zap.Error(err))
	}

	return uint64(count)
}

func QueryIsFavorite(ctx context.Context, uid, vid uint64) bool {
	ok, err := rdb.SIsMember(ctx, fmt.Sprintf("favorite:%d", vid), uid).Result()
	if err != nil {
		zap.L().Error("redis query is follow error", zap.Uint64("user id", uid),
			zap.Uint64("video id", vid), zap.Error(err))
	}

	return ok
}

func CreateFavorite(ctx context.Context, uid, vid uint64) error {
	p := rdb.TxPipeline()
	defer p.Close()

	// 删除目标粉丝列表及当前用户关注列表
	p.SAdd(ctx, fmt.Sprintf("%s:%d", constants.FavoriteKey, vid), uid)
	p.SAdd(ctx, fmt.Sprintf("%s:%d", constants.FavoriteUserKey, uid), vid)
	if _, err := p.Exec(ctx); err != nil {
		p.Discard()
		return err
	}

	return nil
}

func DeleteFavorite(ctx context.Context, uid, vid uint64) error {
	p := rdb.TxPipeline()
	defer p.Close()

	// 删除目标粉丝列表及当前用户关注列表
	p.SRem(ctx, fmt.Sprintf("%s:%d", constants.FavoriteKey, vid), uid)
	p.SRem(ctx, fmt.Sprintf("%s:%d", constants.FavoriteUserKey, uid), vid)
	if _, err := p.Exec(ctx); err != nil {
		p.Discard()
		return err
	}

	return nil
}

func MGetFavoriteVideoIDs(ctx context.Context, uid uint64) (ids []uint64, err error) {
	return GetSetInt64List(ctx, fmt.Sprintf("%s:%d", constants.FavoriteUserKey, uid))
}
