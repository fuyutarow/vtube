package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func StatusCmd(cmd *cobra.Command, args []string) {
	pid := GetMPlayerPID()
	println("pid: ", pid)
}

func GetMPlayerPID() string {
	usr, _ := user.Current()
	pid_fpath := filepath.Join(usr.HomeDir, ".cache/vtube/mplayer.pid")

	file, err := os.Open(pid_fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(b)
}

