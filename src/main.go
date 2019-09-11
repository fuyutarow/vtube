package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	rootCmd := &cobra.Command{Use: "vtube"}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "play",
			Aliases: []string{"p"},
			Run:     PlayCmd,
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "status",
			Short: "Show status",
			Run:   StatusCmd,
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "resume",
			Aliases: []string{"re"},
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("CONT", pid)
			},
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use: "pause",
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("STOP", pid)
			},
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "skip",
			Aliases: []string{"s"},
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("KILL", pid)
			},
		})

	rootCmd.Execute()
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

func SendSig(signal, pid string) {
	cmd := exec.Command("kill", fmt.Sprintf("-%s", signal), pid)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

func StatusCmd(cmd *cobra.Command, args []string) {
	pid := GetMPlayerPID()
	println("pid: ", pid)
}

func PlayCmd(cmd *cobra.Command, args []string) {
	// get webm_cache_dir
	usr, _ := user.Current()
	webm_cache_dir := filepath.Join(usr.HomeDir, ".cache/vtube/webm")
	if err := os.MkdirAll(webm_cache_dir, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := strings.Join(args, " ")
	search_query := strings.Replace(s, " ", "+", -1)
	watchid_list := WatchidListByQuery(search_query)

	// take first
	id := watchid_list[0]
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)

	webm_path := filepath.Join(webm_cache_dir, fmt.Sprintf("%s.webm", id))

	YoutubeDL(webm_path, url).Start()

	if f, err := os.Stat(webm_path); os.IsNotExist(err) || f.IsDir() {
		webmpart_path := filepath.Join(webm_cache_dir, fmt.Sprintf("%s.webm.part", id))
		for {
			if f, err := os.Stat(webmpart_path); os.IsNotExist(err) || f.IsDir() {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}
		MPlayer(webmpart_path)
	} else {
		MPlayer(webm_path)
	}
}

func WatchidListByQuery(query string) []string {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://youtube.com/results?search_query=%s", query))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	watchid_list := []string{}
	doc.Find("a").Each(func(index int, item *goquery.Selection) {
		link, _ := item.Attr("href")
		if strings.Contains(link, "/watch?v=") {
			watchid := strings.Replace(link, "/watch?v=", "", -1)
			watchid_list = append(watchid_list, watchid)
		}
	})

	return watchid_list
}

func MPlayer(fpath string) {
	fmt.Println("[INFO] play", fpath)
	cmd := exec.Command("mplayer", "-noconsolecontrols", "-really-quiet", fpath)
	cmd.Start()

	usr, _ := user.Current()
	mplayer_pid := filepath.Join(usr.HomeDir, ".cache/vtube/mplayer.pid")
	file, err := os.Create(mplayer_pid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	output := fmt.Sprintf("%d", cmd.Process.Pid)
	fmt.Println("pid: ", output)
	file.Write(([]byte)(output))
}

func YoutubeDL(fpath, url string) *exec.Cmd {
	fmt.Println("[INFO] fetch", url)
	cmd := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", fpath, url)
	return cmd
}
