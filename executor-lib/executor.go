package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCode(code string, config ExecutionConfig) (string, error) {
	if err := os.WriteFile(config.Filename, []byte(code), 0644); err != nil {
		return "", fmt.Errorf("file write error: %v", err)
	}

	if config.UseCompiler {
		if out, err := exec.Command(config.CompileCmd[0], config.CompileCmd[1:]...).CombinedOutput(); err != nil {
			return string(out), fmt.Errorf("compilation failed: %v", err)
		}
		_ = exec.Command("cp", config.BinaryName, "/sandbox/"+config.BinaryName).Run()
		_ = exec.Command("chmod", "+x", "/sandbox/"+config.BinaryName).Run()
	} else {
		_ = exec.Command("cp", config.Filename, "/sandbox/"+config.Filename).Run()
	}

	args := []string{
		"--quiet", "--mode", "o", "--chroot", "/sandbox",
		"--user", "9999", "--group", "9999",
		"--rlimit_as", "256", "--time_limit", "3",
		"--disable_proc",
	}

	if config.UseCompiler {
		args = append(args, "--exec_bin", "/"+config.BinaryName, "--", "/"+config.BinaryName)
	} else {
		args = append(args, "--exec_bin", "/"+config.Interpreter, "--", "/"+config.Interpreter, config.Filename)
	}

	cmd := exec.Command("nsjail", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
