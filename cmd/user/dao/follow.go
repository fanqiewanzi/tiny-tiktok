package dao

import (
	"context"
	"fmt"

	"github.com/weirdo0314/tiny-tiktok/pkg/constants"

	"go.uber.org/zap"
)

// QueryFollowerCount 使用redis查询该用户的粉丝列表
func QueryFollowerCount(ctx context.Context, uid uint64) uint64 {
	count, err := rdb.SCard(ctx, fmt.Sprintf("%s:%d", constants.FansKey, uid)).Result()
	if err != nil {
		zap.L().Error("redis query follower count error", zap.Uint64("user id", uid), zap.Error(err))
	}

	return uint64(count)
}

// QueryFollowCount 使用redis查询该用户的关注列表
func QueryFollowCount(ctx context.Context, uid uint64) uint64 {
	count, err := rdb.SCard(ctx, fmt.Sprintf("%s:%d", constants.FollowKey, uid)).Result()
	if err != nil {
		zap.L().Error("redis query follow count error", zap.Uint64("user id", uid), zap.Error(err))
	}

	return uint64(count)
}

// QueryIsFollow 使用redis判断该用户是否已关注
func QueryIsFollow(ctx context.Context, curID, targetID uint64) bool {
	if curID == targetID {
		return true
	}

	ok, err := rdb.SIsMember(ctx, fmt.Sprintf("%s:%d", constants.FansKey, targetID), curID).Result()
	if err != nil {
		zap.L().Error("redis query is follow error", zap.Uint64("user id", curID),
			zap.Uint64("target id", targetID), zap.Error(err))
	}

	return ok
}

func CreateFollow(ctx context.Context, uid, targetID uint64) error {
	p := rdb.TxPipeline()
	defer p.Close()

	// 添加目标粉丝列表及当前用户关注列表
	p.SAdd(ctx, fmt.Sprintf("%s:%d", constants.FansKey, targetID), uid)
	p.SAdd(ctx, fmt.Sprintf("%s:%d", constants.FollowKey, uid), targetID)
	if _, err := p.Exec(ctx); err != nil {
		p.Discard()
		return err
	}

	return nil
}

func DeleteFollow(ctx context.Context, uid, targetID uint64) error {
	p := rdb.TxPipeline()
	defer p.Close()

	// 删除目标粉丝列表及当前用户关注列表
	p.SRem(ctx, fmt.Sprintf("%s:%d", constants.FansKey, targetID), uid)
	p.SRem(ctx, fmt.Sprintf("%s:%d", constants.FollowKey, uid), targetID)
	if _, err := p.Exec(ctx); err != nil {
		p.Discard()
		return err
	}

	return nil
}

func MGetFollowOrFansIDs(ctx context.Context, key string, uid uint64) (ids []uint64, err error) {
	return GetSetUInt64List(ctx, fmt.Sprintf("%s:%d", key, uid))
}
