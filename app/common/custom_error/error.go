package custom_error

import "errors"

var (
	LostUserIdErr = errors.New("用户id缺失")
	SaveErr       = errors.New("保存失败")
)
