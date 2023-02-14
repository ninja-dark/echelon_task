package command

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type Want struct {
	cmd  string
	args []string
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Want
	}{
		{
			name:  "display the disk space",
			input: "df -h",
			want: Want{
				cmd:  "df",
				args: []string{"-h"},
			},
		},
		{
			name:  "ping",
			input: "ping google.com",
			want: Want{
				cmd:  "ping",
				args: []string{"google.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs := NewExecutor().parseArgs(tt.input)
			assert.Equal(t, tt.want.cmd, gotCmd)
			assert.Equal(t, tt.want.args, gotArgs)
		})
	}
}
