package connect_test

import (
	"testing"
	"workspace/pkg/connect"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		url  string
		want bool
	}{
		{name: "基础用例", url: "https://www.baidu.com/", want: true},
		{name: "不存在的url", url: "http://www.thisurldoesnotexistoninternet.com/", want: false},
		{name: "空字符串", url: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := connect.Get(tt.url)
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
