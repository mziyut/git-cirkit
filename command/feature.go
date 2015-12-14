package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func CmdFeature(c *cli.Context) {
	command := os.Args[2]
	branch := make([]byte, 0, 15)
	branch = append(branch, "feature/"...)
	branch = append(branch, os.Args[3]...)

	exec.Command("git", "checkout", "develop").CombinedOutput()

	switch command {
	case "start":
		out, _ := exec.Command("git", "checkout", "-b", string(branch)).CombinedOutput()
		fmt.Println(string(out))
	case "end":
		out, _ := exec.Command("git", "marge", string(branch)).CombinedOutput()
		fmt.Println(string(out))
	}
}
