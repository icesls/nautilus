// +----------------------------------------------------------------------
// | nautilus [ memory implements cache ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package cache

import (
	"sync"
	"time"
)

type MemoryCache struct {
	sync.RWMutex
	data map[string]*data
}

type data struct {
	Data    string
	Expired time.Time
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]*data),
	}
}

// Get 获取缓存内容
func (m *MemoryCache) Get(key string) string {
	if result, ok := m.data[key]; ok {
		if result.Expired.After(time.Now()) {
			m.delKey(key)
			return ""
		}
		return result.Data
	}

	return ""
}

// Set 设置缓存
func (m *MemoryCache) Set(key string, value string, expired time.Duration) error {
	m.Lock()
	defer m.Unlock()

	m.data[key] = &data{
		Data:    value,
		Expired: time.Now().Add(expired),
	}

	return nil
}

// Has 判断key是否存在
func (m *MemoryCache) Has(key string) bool {
	if result, ok := m.data[key]; ok {
		if result.Expired.Before(time.Now()) {
			m.delKey(key)
			return false
		}
		return true
	}

	return false
}

// Del 删除缓存
func (m *MemoryCache) Del(key string) error {
	m.delKey(key)
	return nil
}

func (m *MemoryCache) delKey(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}
