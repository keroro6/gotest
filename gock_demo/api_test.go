package gock_demo

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

//如果不想在测试过程中真正去发送请求或者依赖的外部接口还没有开发完成时，我们可以在单元测试中对依赖的API进行mock
func TestGetResultByAPI(t *testing.T) {
	defer gock.Off()

	// mock 请求外部api时传参x=1返回100, 理解程一个服务的接口，  url是这个， path是这个，     输入参数是1的情况下，   返回的mock数据是100，状态码是200 ，牛啊这个东西
	gock.New("http://your-api.com").Post("/post").MatchType("json").JSON(map[string]int{"x": 1}).Reply(200).JSON(map[string]int{"value": 100})

	//这个函数里面的请求会打到上面 gock.New里
	res := GetResultByAPI(1, 1)
	assert.Equal(t, res, 101)

	gock.New("http://your-api.com").Post("/post").MatchType("json").JSON(map[string]int{"x": 2}).Reply(200).JSON(map[string]int{"value": 200})

	res = GetResultByAPI(2, 2)
	assert.Equal(t, res, 202)

	assert.True(t, gock.IsDone())
}
