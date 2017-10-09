package cmd

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func filter(history string) (commands string, err error) {
	var buf bytes.Buffer
	err = run("peco", strings.NewReader(history), &buf)
	if err != nil {
		return "", nil
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	return lines[0], nil
}
