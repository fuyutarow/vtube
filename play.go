package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"os/exec"
	"strings"
	// "time"
)

func main() {
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

	fmt.Println(url)
	proc := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", "%(id)s.%(ext)s", url)
	proc.Start()

	webm := fmt.Sprintf("%s.webm", id)
	if f, err := os.Stat(webm); os.IsNotExist(err) || f.IsDir() {
		webm_part := fmt.Sprintf("%s.webm.part", id)
		fmt.Println(webm_part)
		MPlayer(webm_part)
	} else {
		fmt.Println(webm)
		MPlayer(webm)
	}

	proc.Wait()
}

func MPlayer(fpath string) {
	out, err := exec.Command("mplayer", "-noconsolecontrols", "-really-quiet", fpath).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func YoutubeDL(id string) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)

	fmt.Println(url)
	proc := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", "%(id)s.%(ext)s", url)
	proc.Run()

	// out, err := exec.Command("youtube-dl", "-f", "worstaudio[ext=webm]", "-o", "%(id)s.%(ext)s", url).Output()
	// if err != nil {
	//     fmt.Println(err)
	//     os.Exit(1)
	// }
	// fmt.Println(string(out))
}
