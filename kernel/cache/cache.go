// +----------------------------------------------------------------------
// | nautilus [ cache instance ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package cache

import (
	"encoding/json"
	"github.com/icesls/nautilus/kernel/logger"
	"sync"
	"time"
)

var (
	once  sync.Once
	Cache *cache
)

type cache struct {
	Store IStore
}

// Init 缓存驱动从外部传进来
func Init(store IStore) {
	once.Do(func() {
		Cache = &cache{
			Store: store,
		}
	})
}

// Set 设置缓存
func Set(key string, value any, expire time.Duration) {
	val, err := json.Marshal(value)
	logger.LogIf(err)
	Cache.Store.Set(key, string(val), expire)
}

// Get 获取缓存内容
func Get(key string) any {
	var result any
	value := Cache.Store.Get(key)
	err := json.Unmarshal([]byte(value), &result)
	logger.LogIf(err)

	return result
}

// Has 判断是否存在
func Has(key string) bool {
	return Cache.Store.Has(key)
}

// Del 删除缓存内容
func Del(key string) bool {
	err := Cache.Store.Del(key)
	if err != nil {
		logger.LogIf(err)
		return false
	}

	return true
}
