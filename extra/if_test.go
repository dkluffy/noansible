package extra

import "testing"

func TestIfTester(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{

			name: "1",
			args: args{s: "aa;= ;a"},
			want: false,
		},
		{

			name: "2",
			args: args{s: "a;=  ;a"},
			want: true,
		},
		{

			name: "3",
			args: args{s: "123;=;a"},
			want: false,
		},
		{

			name: "4",
			args: args{s: "123;>   ;1 1"},
			want: true,
		},
		{

			name: "5",
			args: args{s: "aa"},
			want: true,
		},
		{

			name: "6",
			args: args{s: "    "},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfTester(tt.args.s); got != tt.want {
				t.Errorf("IfTester() = %v, want %v", got, tt.want)
			}
		})
	}
}
