package monkeytest

import (
	"fmt"
	"gotest/monkeytest/varys"
)

func MyFunc(uid int64) string {
	u, err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s\n", u.Name)
}
