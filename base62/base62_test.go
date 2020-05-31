package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeDeocde(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "encodes simple string",
			args: args{
				src: "simple",
			},
		},
		{
			name: "enocdes numbers",
			args: args{
				src: "52341",
			},
		},
		{
			name: "hyphenated number string number",
			args: args{
				src: "abhishek-123",
			},
		},
		{
			name: "long alphanumeric hypen",
			args: args{
				src: "QR-13959333-1590488716162411758",
			},
		},
		{
			name: "nihongo non ascii",
			args:args{
				src:"日本語",
			},
		},
		{
			name: "ascii mixed with non ascii ",
			args:args{
				src:"ab日本語-123de",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.args.src, Decode(Encode(tt.args.src)))
		})
	}
}
