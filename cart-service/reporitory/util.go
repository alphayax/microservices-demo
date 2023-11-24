package reporitory

import (
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Initialize(redisUri string) error {
	opt, err := redis.ParseURL(redisUri)
	if err != nil {
		return errors.Wrap(err, "Unable to initialize the redis connection")
	}

	client = redis.NewClient(opt)
	return nil
}
