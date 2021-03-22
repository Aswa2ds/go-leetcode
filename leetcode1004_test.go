package leetcode

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{
				A: []int{1,0,1,1,1,0,0,1,1},
				K: 2,
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestOnes(tt.args.A, tt.args.K); gotAns != tt.wantAns {
				t.Errorf("longestOnes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
