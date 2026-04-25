package executor

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type BaseExecutor struct{}

func (b BaseExecutor) runLocalCommand(
	code string,
	fileName string,
	command []string,
) (string, error) {

	// Create temp dir
	tempDir, err := os.MkdirTemp("", "exec-*")
	if err != nil {
		return "Failed to create temp dir", err
	}
	defer os.RemoveAll(tempDir)

	// Write file
	filePath := filepath.Join(tempDir, fileName)
	err = os.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return "Failed to write file", err
	}

	cmdArgs := make([]string, len(command))
	copy(cmdArgs, command)

	for i, arg := range cmdArgs {
		arg = strings.ReplaceAll(arg, "{file}", filePath)
		arg = strings.ReplaceAll(arg, "{dir}", tempDir)
		cmdArgs[i] = arg
	}

	// Execute command
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = tempDir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	var output []byte
	var runErr error

	done := make(chan struct{})

	go func() {
		output, runErr = cmd.CombinedOutput()
		close(done)
	}()

	select {
	case <-done:
		// finished normally
		if runErr != nil {
			return string(output) + "\nError: " + runErr.Error(), nil
		}
		return string(output), nil

	case <-time.After(2 * time.Second):
		// timeout reached → kill process group
		if cmd.Process != nil {
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return "Execution timed out", nil
	}
}
