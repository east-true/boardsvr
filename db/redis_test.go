package db_test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestExampleClient(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Error(err)
		return
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("key %s", val)

	val2 := rdb.Del(ctx, "key")
	if val2.Err() != nil {
		t.Error(val2.Err())
		return
	}
}
