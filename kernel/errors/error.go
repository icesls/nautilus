// +----------------------------------------------------------------------
// | nautilus [ errors ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package errors

// Errors 微信返回的错误字段信息
type Errors struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}
