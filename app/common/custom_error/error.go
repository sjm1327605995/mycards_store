package custom_error

import "errors"

var (
	LostUserIdErr = errors.New("用户id缺失")
	SaveErr       = errors.New("保存失败")
	NotFoundErr   = errors.New("记录不存在")
	FoundErr      = errors.New("查询失败")
	DeleteErr     = errors.New("删除失败")
)
