package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xddbom/rt-chat/db"
)

func TestRedisSetup(t *testing.T) {
	rdb := db.RedisSetup()
	defer rdb.Close()

	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	assert.NoError(t, err, "Redis must be available")
}