package cache

import (
	"fmt"
	"passport.xinfos.com/driver"
	"passport.xinfos.com/internal/model"
	"time"

	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	module             = "cache.user"
	cacheUserDetailKey = "cache:userdetail:%d"
)

//UserCache - user cache
type UserCache struct {
	Rds *redis.Client
}

func NewUserCache() *UserCache {
	return &UserCache{
		Rds: driver.Rds,
	}
}

func (u *UserCache) Get(id uint64) (user *model.User) {
	key := fmt.Sprintf(cacheUserDetailKey, id)
	v, err := u.Rds.Get(key).Bytes()
	if err != nil {
		return nil
	}
	err = msgpack.Unmarshal(v, &user)
	if err != nil {
		return nil
	}
	return user
}

func (u *UserCache) Set(user *model.User) bool {
	k := fmt.Sprintf(cacheUserDetailKey, user.ID)
	v, err := msgpack.Marshal(user)
	if err != nil {
		return false
	}
	isOk, err := u.Rds.SetNX(k, v, 100*time.Second).Result()
	if err != nil {
		return false
	}
	return isOk
}
