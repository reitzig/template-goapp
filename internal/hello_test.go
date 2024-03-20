package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assertThat := assert.New(t)

	type args struct {
		dear   bool
		hailee string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{name: "polite", args: args{dear: true, hailee: "Kevin"}, expected: "Hello, dear Kevin!"},
		{name: "impolite", args: args{dear: false, hailee: "Stuart"}, expected: "Hello, Stuart!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			config := HelloConfig{Dear: ConfigValue[bool]{Source: "test", Value: tt.args.dear}}
			assertThat.Equal(tt.expected, Hello(config, tt.args.hailee))
		})
	}
}
