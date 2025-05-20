package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCode(code string, config ExecutionConfig) (string, error) {
	if err := os.WriteFile(config.Filename, []byte(code), 0644); err != nil {
		return "", fmt.Errorf("file write error: %v", err)
	}

	if config.UseCompiler {
		compileCmd := exec.Command(config.CompileCmd[0], config.CompileCmd[1:]...)
		compileCmd.Dir, _ = filepath.Abs(".")
		if out, err := compileCmd.CombinedOutput(); err != nil {
			return string(out), fmt.Errorf("compilation failed: %v", err)
		}
	} else {
		_ = os.Chmod(config.Filename, 0755)
	}

	var cmd *exec.Cmd
	if len(config.RunCmd) > 0 {
		cmd = exec.Command(config.RunCmd[0], config.RunCmd[1:]...)
	} else if config.UseCompiler {
		cmd = exec.Command("./" + config.BinaryName)
	} else {
		cmd = exec.Command(config.Interpreter, config.Filename)
	}

	cmd.Dir, _ = filepath.Abs(".")

	out, err := cmd.CombinedOutput()
	return string(out), err
}
