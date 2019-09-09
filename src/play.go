package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	usr, _ := user.Current()
	webm_cache_dir := filepath.Join(usr.HomeDir, ".cache/vtube/webm")
	if err := os.MkdirAll(webm_cache_dir, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("[INFO] mkdir", webm_cache_dir)
	}

	s := strings.Join(os.Args[1:], " ")
	q := strings.Replace(s, " ", "+", -1)
	doc, err := goquery.NewDocument(fmt.Sprintf("https://youtube.com/results?search_query=%s", q))
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

	// take first
	id := watchid_list[0]
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)

	fmt.Println("[INFO] fetch", url)

	webm_path := filepath.Join(webm_cache_dir, fmt.Sprintf("%s.webm", id))

	// $ youtube-dl -f "worstaudio[ext=webm]" -o "%(id)s.%(ext)s" url
	proc := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", webm_path, url)
	proc.Start()
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

	proc.Wait()
}

func MPlayer(fpath string) {
	fmt.Println("[INFO] play", fpath)
	out, err := exec.Command("mplayer", "-noconsolecontrols", "-really-quiet", fpath).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
