package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test: end with '/'",
			args: args{path: "/home/"},
			want: "/home",
		},
		{
			name: "test: /../",
			args: args{path: "/../"},
			want: "/",
		},
		{
			name: "test: contains multi /",
			args: args{path: "/home//aswa2ds"},
			want: "/home/aswa2ds",
		},
		{
			name: "test: normal case",
			args: args{path: "/a/./b/../../c"},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
