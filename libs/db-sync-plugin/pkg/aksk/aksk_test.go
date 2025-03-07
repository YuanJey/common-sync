package aksk

import "testing"

func TestGetAkSK(t *testing.T) {
	type args struct {
		skPart string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetAkSK(tt.args.skPart)
			if got != tt.want {
				t.Errorf("GetAkSK() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAkSK() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
