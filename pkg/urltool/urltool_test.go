package urltool_test

import (
	"testing"
	"workspace/pkg/urltool"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		targetUrl string
		want      string
		wantErr   bool
	}{
		{name: "基本示例", targetUrl: "https://www.liwenzhou.com/posts/Go/unit-test-1/", want: "unit-test-1", wantErr: false},
		{name: "相对路径url示例", targetUrl: "/xxx/1233/", want: "", wantErr: true},
		{name: "空字符串", targetUrl: "", want: "", wantErr: true},
		{name: "带query的url", targetUrl: "https://www.liwenzhou.com/posts/Go/unit-test-1/?a=1&b=2", want: "unit-test-1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urltool.GetBasePath(tt.targetUrl)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBasePath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBasePath() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
