package support

import (
	"fmt"
	go_stringable "github.com/hyperf/go-stringable"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"sync/atomic"
	"time"
)

type IdGenerator struct {
	Incrementer     IncrementerInterface
	SuffixMaxLength uint64
}

type DefaultIncrementer struct {
	Id uint64
}

func (d *DefaultIncrementer) Incr() (uint64, error) {
	return atomic.AddUint64(&d.Id, 1), nil
}

type RedisIncrementer struct {
	Rds *redis.Redis
	Key string
}

func (r *RedisIncrementer) Incr() (uint64, error) {
	res, err := r.Rds.Incr(r.Key)
	if err != nil {
		return 0, err
	}

	return uint64(res), nil
}

type IncrementerInterface interface {
	Incr() (uint64, error)
}

type IdGeneratorInterface interface {
	Generate() (uint64, error)
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{&DefaultIncrementer{Id: 0}, 4}
}

func NewRedisIdGenerator(rds *redis.Redis, key string) *IdGenerator {
	return &IdGenerator{&RedisIncrementer{Rds: rds, Key: key}, 4}
}

func (g *IdGenerator) Generate() (uint64, error) {
	id, err := g.Incrementer.Incr()
	if err != nil {
		return 0, err
	}

	result := fmt.Sprintf("%d%s", time.Now().Unix(), g.Suffix(id))
	return strconv.ParseUint(result, 10, 64)
}

func (g *IdGenerator) Suffix(id uint64) string {
	maxValue := 10
	exp := g.SuffixMaxLength - 1
	for exp > 0 {
		maxValue *= 10
		exp--
	}

	return go_stringable.StrPadLeft(strconv.FormatUint(id%uint64(maxValue), 10), '0', 4)
}
