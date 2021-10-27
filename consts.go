package selector

import "fmt"

var (
	errUnknownAuthWay = fmt.Errorf("未知的权限路径")
	errUnknownRole    = fmt.Errorf("未知的角色")
	errWrongAuthType  = fmt.Errorf("未知的权限类型")
	errWrongState     = fmt.Errorf("未知的权限状态")
)
