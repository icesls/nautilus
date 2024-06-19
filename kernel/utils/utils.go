// +----------------------------------------------------------------------
// | nautilus [ utils ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package utils

import "encoding/json"

func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

func ToJsonString(data interface{}) string {
	return string(toJson(data))
}

func ToJsonBytes(data interface{}) []byte {
	return toJson(data)
}

func toJson(data interface{}) []byte {
	if data == nil {
		return nil
	}
	marshal, _ := json.Marshal(data)
	return marshal
}
