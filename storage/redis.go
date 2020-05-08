package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mattheath/base62"
	"shortlink/utils"
	"time"
)

const (
	URLIDKEY           = "next.url.id"
	ShortlinkKey       = "shortlink:%s:url"
	URLHashKey         = "urlHash:%s:url"
	ShortlinkDetailKey = "shortlink:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}
type URLDetail struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"create_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

func NewRedis(addr, pwd string, db int) *RedisCli {
	client := redis.NewClient(&redis.Options{Addr: addr, Password: pwd, DB: db})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{client}
}
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	sha1 := utils.ToSha1(url)
	result, err := r.Cli.Get(fmt.Sprintf(URLHashKey, sha1)).Result()
	if err == redis.Nil {

	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}
	err = r.Cli.Incr(fmt.Sprintf(URLIDKEY)).Err()
	if err != nil {
		return "", err
	}
	id, err := r.Cli.Get(URLIDKEY).Int64()
	if err != nil {
		return "", err
	}
	eid := base62.EncodeInt64(id)
	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	marshal, err := json.Marshal(&URLDetail{
		URL:                 url,
		CreateAt:            time.Now().String(),
		ExpirationInMinutes: time.Duration(exp),
	})
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), marshal, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", nil
	}
	return eid, nil
}
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	result, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", errors.New("unknow short URL")
	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
func (r *RedisCli) Unshorten(eid string) (string, error) {
	result, err := r.Cli.Get(fmt.Sprintf(ShortlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", errors.New("unknow short URL")
	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}

}
