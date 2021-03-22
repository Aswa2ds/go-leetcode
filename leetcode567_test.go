package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s1: "abc",
				s2: "acbd",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s1: "abc",
				s2: "ab",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s1: "acb",
				s2: "fafdsafdabcsdfsdf",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
