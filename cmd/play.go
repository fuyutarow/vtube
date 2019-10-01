package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/PuerkitoBio/goquery"
	"os/exec"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

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

	var webm_path = filepath.Join(webm_cache_dir, fmt.Sprintf("%s.webm", id))

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
		webm_path = webmpart_path
	}
	MPlayer(webm_path)

	fpath := filepath.Join(usr.HomeDir, ".cache/vtube/playing")
	file, err := os.Create(fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	output := webm_path
	fmt.Println("[INFO] pid:", output)
	file.Write(([]byte)(output))
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
	fmt.Println("[INFO] pid:", output)
	file.Write(([]byte)(output))
}

func YoutubeDL(fpath, url string) *exec.Cmd {
	fmt.Println("[INFO] fetch", url)
	cmd := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", fpath, url)
	return cmd
}