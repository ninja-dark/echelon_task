package command

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

const (
	testStdoutValue = "test fun value!"
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

/*
var mockedExitStatus = 0
var mockedStdout string


func fakeExecCommand(command string, args ...string) *exec.Cmd {
    cs := []string{"-test.run=TestExecCommandHelper", "--", command}
    cs = append(cs, args...)
    cmd := exec.Command(os.Args[0], cs...)
    es := strconv.Itoa(mockedExitStatus)
    cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1",
        "STDOUT=" + mockedStdout,
        "EXIT_STATUS=" + es}
    return cmd
}

func TestExecCommandHelper(t *testing.T) {
    if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
        return
    }

    // println("Mocked stdout:", os.Getenv("STDOUT"))
    fmt.Fprintf(os.Stdout, os.Getenv("STDOUT"))
    i, _ := strconv.Atoi(os.Getenv("EXIT_STATUS"))
    os.Exit(i)
}
func TestPrintDate(t *testing.T) {
    mockedExitStatus = 1
    mockedStdout = "Hello Worl"
    cmd:= exec.Command(fakeExecCommand)
    defer func() { cmd = exec.Command }()
    expDate := "Hello Worl"

    out, _ := printDate()
    if string(out) != expDate {
        t.Errorf("Expected %q, got %q", expDate, string(out))
    }
}
*/
