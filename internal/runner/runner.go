package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCommands(commands []string) error {

	if len(commands) == 0 {
		return nil
	}
	// " && " で連結
	joined := strings.Join(commands, " && ")
	fmt.Println("> " + joined)
	cmd := exec.Command("sh", "-c", joined)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()

}