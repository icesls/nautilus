// +----------------------------------------------------------------------
// | nautilus [ cache interface ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package cache

import "time"

type IStore interface {
	Set(key, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Del(key string) error
}
