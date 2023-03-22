package redis_test

import (
	cartRedis "api/src/cart/repository/redis"
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type CartIntegrationTestSuite struct {
	suite.Suite
	R cartRedis.CartRedisRepo
}

func (s *CartIntegrationTestSuite) SetupTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s.R = cartRedis.CartRedisRepo{Conn: rdb}
}

func (s *CartIntegrationTestSuite) TearDownTest() {
	ctx := context.Background()
	result := s.R.Conn.FlushDB(ctx)
	s.Assertions.NoError(result.Err())
}

func TestCartRedisRepository(t *testing.T) {
	suite.Run(t, new(CartIntegrationTestSuite))
}
