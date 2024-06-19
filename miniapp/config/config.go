// +----------------------------------------------------------------------
// | nautilus [ officialAccount config ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package config

import "github.com/icesls/nautilus/kernel/cache"

type Config struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Cache     cache.IStore
}
