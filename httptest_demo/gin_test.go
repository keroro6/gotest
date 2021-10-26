package httptest_demo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//对接口做测试
func Test_helloHandler(t *testing.T) {
	// 定义两个测试用例
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{name: "base case", param: `{"name": "liwenzhou"}`, expect: "hello liwenzhou"},
		{name: "bad case", param: "", expect: "we need a name"},
	}

	r := SetupRouter()

	for _, tt := range tests {
		req := httptest.NewRequest("POST", "/hello", strings.NewReader(tt.param))

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)
		assert.Equal(t, tt.expect, resp["msg"])
	}

}
