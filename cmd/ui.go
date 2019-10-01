package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func GetPlayingPath() string {
	usr, _ := user.Current()
	pid_fpath := filepath.Join(usr.HomeDir, ".cache/vtube/playing")

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

func RunUI() {
	webm_path := GetPlayingPath()
	if _, err := exec.Command("open", webm_path).Output(); err != nil {
	} else {
		fmt.Println(webm_path)
		os.Exit(0)
	}

	webmpart_path := strings.TrimRight(webm_path, ".part")
	if _, err := exec.Command("open", webmpart_path).Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(webmpart_path)
		os.Exit(0)
	}

}

func UiCmd(cmd *cobra.Command, args []string) {
	RunUI()
}
