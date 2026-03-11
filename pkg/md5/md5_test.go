package md5_test

import (
	"testing"
	"workspace/pkg/md5"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data []byte
		want string
	}{
		{name: "基本示例", data: []byte("hello world"), want: "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{name: "空字符串", data: []byte(""), want: "d41d8cd98f00b204e9800998ecf8427e"},
		{name: "特殊字符", data: []byte("!@#$%^&*()"), want: "05b28d17a7b6e7024b6e5d8cc43a8bf7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := md5.Sum(tt.data)
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
