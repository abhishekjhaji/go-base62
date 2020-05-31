package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toBase62(t *testing.T) {
	type args struct {
		num uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "19 digits long",
			args: args{
				num: 7017285971513730411,
			},
			want: "8MNF8h7B3D1",
		},
		{
			name: "uint64 max value",
			args: args{
				num: 18446744073709551615,
			},
			want: "LygHa16AHYF",
		},
		{
			name: "single digit",
			args: args{
				num: 7,
			},
			want: "7",
		},
		{
			name: "5 digits",
			args: args{
				num: 30410,
			},
			want: "7uU",
		},
		{
			name: "9 digits",
			args: args{
				num: 859715137,
			},
			want: "wBHBB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBase62(tt.args.num); got != tt.want {
				t.Errorf("toBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromBase62(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "19 digits long",
			args: args{
				encoded: "8MNF8h7B3D1",
			},
			want: 7017285971513730411,
		},
		{
			name: "uint64 max value",
			args: args{
				encoded: "LygHa16AHYF",
			},
			want: 18446744073709551615,
		},
		{
			name: "single digit",
			args: args{
				encoded: "7",
			},
			want: 7,
		},
		{
			name: "5 digits",
			args: args{
				encoded: "7uU",
			},
			want: 30410,
		},
		{
			name: "9 digits",
			args: args{
				encoded: "wBHBB",
			},
			want: 859715137,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromBase62(tt.args.encoded)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
