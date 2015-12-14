package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os/exec"
)

func CmdInit(c *cli.Context) {
	out, _ := exec.Command("git", "checkout", "-b", "develop").CombinedOutput()
	fmt.Println(string(out))
}
