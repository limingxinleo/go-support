package support

import (
	"fmt"
	go_stringable "github.com/hyperf/go-stringable"
	"strconv"
	"sync/atomic"
	"time"
)

type IdGenerator struct {
	Id              uint64
	SuffixMaxLength uint64
}

type IdGeneratorInterface interface {
	Generate() (uint64, error)
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{0, 4}
}

func (g *IdGenerator) Generate() (uint64, error) {
	result := fmt.Sprintf("%d%s", time.Now().Unix(), g.Suffix())
	return strconv.ParseUint(result, 10, 64)
}

func (g *IdGenerator) Suffix() string {
	maxValue := 10
	exp := g.SuffixMaxLength - 1
	for exp > 0 {
		maxValue *= 10
		exp--
	}

	id := atomic.AddUint64(&g.Id, 1)

	return go_stringable.StrPadLeft(strconv.FormatUint(id%uint64(maxValue), 10), '0', 4)
}
