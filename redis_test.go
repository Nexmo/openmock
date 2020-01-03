package openmock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	om := &OpenMock{}
	om.SetRedis()

	t.Run("get non-exists data", func(t *testing.T) {
		v, err := redisHandleReply(om.redis.Do("get", "non-exist"))
		assert.NoError(t, err)
		assert.Empty(t, v)
	})

	t.Run("set data", func(t *testing.T) {
		v, err := redisHandleReply(om.redis.Do("set", "hello", "world"))
		assert.NoError(t, err)
		assert.NotEmpty(t, v)
	})

	t.Run("get data", func(t *testing.T) {
		v, err := redisHandleReply(om.redis.Do("get", "hello"))
		assert.NoError(t, err)
		assert.Equal(t, "world", v)
	})

	t.Run("rpush set array data", func(t *testing.T) {
		_, err := redisHandleReply(om.redis.Do("rpush", "k1", "v1"))
		assert.NoError(t, err)

		v, err := redisHandleReply(om.redis.Do("rpush", "k1", "v2"))
		assert.NoError(t, err)
		assert.NotEmpty(t, v)
	})

	t.Run("rpush get array data", func(t *testing.T) {
		v, err := redisHandleReply(om.redis.Do("lrange", "k1", 0, -1))
		assert.NoError(t, err)
		assert.Equal(t, "v1;;v2", v)
	})
}

func TestRedisDo(t *testing.T) {
	om := &OpenMock{
		RedisType: "",
	}
	r := redisDo(om)

	t.Run("get non-exists", func(t *testing.T) {
		v := r("GET", "non-exists")
		assert.Empty(t, v)
	})

	t.Run("set and then get", func(t *testing.T) {
		r("SET", "hello", "456")
		v := r("GET", "hello")
		assert.Equal(t, "456", v)
	})

	t.Run("errors on template store args", func(t *testing.T) {
		key := redisTemplatesStore + "_asdf123"
		_, err := om.redis.Do("HSET", key, "123", "stuff")
		assert.NoError(t, err)

		v := r("SET", key, "123", "things")
		assert.Error(t, v.(error))

		_, err = om.redis.Do("DEL", key)
		assert.NoError(t, err)
	})
}
