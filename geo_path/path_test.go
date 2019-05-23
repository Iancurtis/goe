package geo_path

import "testing"

func TestFileBaseName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{"a.b"},
			want: "a",
		},
		{
			args: args{"a"},
			want: "a",
		},
		{
			args: args{"a.b.c"},
			want: "a.b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileBaseName(tt.args.fileName); got != tt.want {
				t.Errorf("FileBaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}
