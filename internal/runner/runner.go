package runner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func RunCommands(commands []string) error {

	if len(commands) == 0 {
		return nil
	}
	// " && " で連結
	joined := strings.Join(commands, " && ")
	fmt.Println("> " + joined)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", joined)
	} else {
		cmd = exec.Command("sh", "-c", joined)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()

}
