package goconveytest

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

//func TestSplit(t *testing.T) {
//	convey.Convey("基础用例", t, func() {
//		var (
//			s      = "a:b:c"
//			sep    = ":"
//			expect = []string{"a", "b", "c"}
//		)
//
//		got := Split(s, sep)
//		convey.So(got, convey.ShouldResemble, expect) //断言
//	})
//
//	convey.Convey("不包含分隔符用例", t, func() {
//		var (
//			s      = "a:b:c"
//			sep    = "|"
//			expect = []string{"a:b:c"}
//		)
//		got := Split(s, sep)
//		convey.So(got, convey.ShouldResemble, expect)
//	})
//}

//goconvey还支持在单元测试中根据需要嵌套调用，比如：

func TestSplit(t *testing.T) {
	// 只需要在顶层的Convey调用时传入t
	convey.Convey("分隔符在开头或结尾用例", t, func() {
		tt := []struct {
			name   string
			s      string
			sep    string
			expect []string
		}{
			{"分隔符在开头", "*1*2*3", "*", []string{"", "1", "2", "3"}},
			{"分隔符在结尾", "1+2+3+", "+", []string{"1", "2", "3", ""}},
		}

		for _, tc := range tt {
			convey.Convey(tc.name, func() { // 嵌套调用Convey
				got := Split(tc.s, tc.sep)
				convey.So(got, convey.ShouldResemble, tc.expect) // convey.ShouldResemble 用于数组、切片、map和结构体相等
			})
		}
	})
}
