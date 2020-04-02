package src

import (
	"errors"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"github.com/ArkNX/ark-go/plugin/redisPlugin/redisModule"
	"github.com/go-redis/redis"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

// ErrInvalidRedisAddr describes error of invalid redis address
var ErrInvalidRedisAddr = errors.New("invalid redis address")

func init() {
	t := reflect.TypeOf((*CRedisModule)(nil))
	if !t.Implements(reflect.TypeOf((*redisModule.IRedisModule)(nil)).Elem()) {
		log.Fatal("AFIRedisModule is not implemented by CRedisModule")
	}

	redisModule.ModuleType = t.Elem()
	redisModule.ModuleName = filepath.Join(redisModule.ModuleType.PkgPath(), redisModule.ModuleType.Name())
	redisModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CRedisModule{}).Update).Pointer()).Name()
}

type CRedisModule struct {
	ark.Module
	// other module
	log logModule.ILogModule
	// other data
	conn redis.Cmdable
}

func (redisModule *CRedisModule) Init() error {
	m := redisModule.GetPluginManager().FindModule(logModule.ModuleName)
	logModule, ok := m.(logModule.ILogModule)
	if !ok {
		log.Fatal("failed to get log module in httpServer module")
	}
	redisModule.log = logModule
	return nil
}

func (redisModule *CRedisModule) Connect(addrs []string, password string, poolSize int) error {
	if len(addrs) == 0 {
		return ErrInvalidRedisAddr
	}

	var conn redis.Cmdable
	if len(addrs) > 1 {
		conn = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    addrs,
			Password: password,
			PoolSize: poolSize,
		})
	} else {
		conn = redis.NewClient(&redis.Options{
			Addr:     addrs[0],
			Password: password,
			PoolSize: poolSize,
		})
	}

	if _, err := conn.Ping().Result(); err != nil {
		redisModule.log.GetLogger().WithFields(map[string]interface{}{
			"redisAddr": addrs,
		}).Error("failed to ping redis during connection")
		return err
	}

	redisModule.conn = conn

	return nil

}

func (redisModule *CRedisModule) GetConn() redis.Cmdable {
	return redisModule.conn
}

// --------------- some basic cmd ---------------
func (redisModule *CRedisModule) Get(key string) (string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return "", err
	}

	return redisModule.conn.Get(key).Result()
}

func (redisModule *CRedisModule) Set(key string, value interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	return redisModule.conn.Set(key, value, expiration).Err()
}

func (redisModule *CRedisModule) INCR(key string) (int64, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return 0, err
	}

	return redisModule.conn.Incr(key).Result()
}

func (redisModule *CRedisModule) INCRBy(key string, value int64) (int64, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return 0, err
	}

	return redisModule.conn.IncrBy(key, value).Result()
}

func (redisModule *CRedisModule) HSet(key, field string, value interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	if err := redisModule.conn.HSet(key, field, value).Err(); err != nil {
		return err
	}

	if err := redisModule.conn.Expire(key, expiration).Err(); err != nil {
		redisModule.conn.Del(key)
		return err
	}

	return nil
}

func (redisModule *CRedisModule) HMSet(key string, fields map[string]interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	if err := redisModule.conn.HMSet(key, fields).Err(); err != nil {
		return err
	}

	if err := redisModule.conn.Expire(key, expiration).Err(); err != nil {
		redisModule.conn.Del(key)
		return err
	}

	return nil
}

func (redisModule *CRedisModule) HGet(key, field string) (string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return "", err
	}

	return redisModule.conn.HGet(key, field).Result()
}

func (redisModule *CRedisModule) HGetAll(key string) (map[string]string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return nil, err
	}

	return redisModule.conn.HGetAll(key).Result()
}

func (redisModule *CRedisModule) Del(keys ...string) {
	redisModule.conn.Del(keys...)
}
