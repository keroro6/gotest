package monkeytest

import (
	"bou.ke/monkey"
	"gotest/monkeytest/varys"
	"strings"
	"testing"
)

func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &varys.UserInfo{Name: "liwenzhou"}, nil
	monkey.Patch(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "liwenzhou") {
		t.Fatal()
	}
}
