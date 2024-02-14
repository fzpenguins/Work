package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"videoweb/database/DB/model"
)

func GetCommentLikeCountKey(cid int64) string {
	return fmt.Sprintf("like_count/comment:%s", strconv.Itoa(int(cid)))
}

func GetLikeFromUser(uid int64) string {
	return fmt.Sprintf("%s/like_count/comment", strconv.Itoa(int(uid))) //传入点赞的人的uid
}

func AddCommentLikeCount(ctx context.Context, uid int64, comment *model.Comment) {

	if RedisClient.ZScore(ctx, GetLikeFromUser(uid), strconv.FormatInt(comment.Cid, 10)).Val() != 1 {
		//还要有一个所有人可见的页面
		RedisClient.Incr(ctx, GetVideoLikeCountKey(comment.Cid))
		RedisClient.ZIncrBy(ctx, "comment:like_count", 1, strconv.Itoa(int(comment.Cid)))
	}

	RedisClient.ZAdd(ctx, GetLikeFromUser(uid), redis.Z{Member: comment.Cid, Score: 1}) //comment.Cid, 1) //自己的页面
}

func CommentLikeCount(ctx context.Context, cid int64) int64 {
	countStr, _ := RedisClient.Get(ctx, GetCommentLikeCountKey(cid)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count

}

func DecrCommentLikeCount(ctx context.Context, uid int64, comment *model.Comment) {
	if RedisClient.ZScore(ctx, GetLikeFromUser(uid), strconv.FormatInt(comment.Cid, 10)).Val() != 0 {
		//还要有一个所有人可见的页面
		RedisClient.Decr(ctx, GetVideoLikeCountKey(comment.Cid))
		RedisClient.ZIncrBy(ctx, "comment:like_count", -1, strconv.Itoa(int(comment.Cid)))
	}
	RedisClient.ZAdd(ctx, GetLikeFromUser(uid), redis.Z{Member: comment.Cid, Score: 0})
}
