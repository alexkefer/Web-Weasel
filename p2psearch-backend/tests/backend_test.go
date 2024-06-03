package tests

import (
	"io"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestBackend(t *testing.T) {
	cmd := exec.Command("go", "run", "..")
	Start(t, cmd)

	time.Sleep(time.Second)

	Kill(t, cmd)
}

func Start(t *testing.T, cmd *exec.Cmd) {
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		t.Fatalf("cound not start backend: %s\n", err)
	}
}

func Kill(t *testing.T, cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		t.Fatalf("cound not kill backend: %s\n", err)
	}
}

func Pipe(pipe io.Reader) {
	buffer := make([]byte, 256)
	for {
		if n, err := pipe.Read(buffer); err != nil {
			return
		} else if n == 0 {
			return
		} else {
			os.Stdout.Write(buffer)
		}
	}

}
