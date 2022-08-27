package post

import (
	"context"
	"fmt"

	"github.com/lily-lee/blog-server/config"
	"github.com/sirupsen/logrus"
)

func CommentKey(postId uint64) string {
	return fmt.Sprintf("posts:comment:post_id:%d", postId)
}

func LikeKey(postId uint64) string {
	return fmt.Sprintf("posts:like:post_id:%d", postId)
}

func SaveNum(key string, add bool, postId uint64, numFunc func(postId uint64) int64) {
	rc := config.RC
	i, err := rc.Exists(context.Background(), key).Result()
	if err != nil {
		logrus.Errorf("redis exists failed. key:%s, post_id:%d, err:%v", key, postId, err)
		return
	}

	// key not exists
	if i == 0 {
		numFunc(postId)
		return
	}

	// incr
	if add {
		_, err = rc.Incr(context.Background(), key).Result()
		if err != nil {
			logrus.Errorf("incr num failed. key:%s, post_id:%d, err:%v", key, postId, err)
		}
		return
	}

	// decr
	num, err := rc.Get(context.Background(), key).Int64()
	if num <= 0 {
		num = num * -1
	} else {
		num = -1
	}

	_, err = rc.IncrBy(context.Background(), key, num).Result()
	if err != nil {
		logrus.Errorf("incrby num failed. key:%s, post_id:%d, err:%v", key, postId, err)
	}
}
