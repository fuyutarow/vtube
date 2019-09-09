
import (
    os
    http
)

fn main () {

    if os.args.len < 1 {
        println('required args')
        return
    }

    query := os.args.right(1).join(' ').replace(' ','+')
    html := http.get_text(
        'https://www.youtube.com/results?search_query=${query}'
    )

    // take first url
    mut pos := 0
    pos = html.index_after('/watch?v=', pos + 1)
    end := html.index_after('"', pos)

    id := html.substr(pos,end).replace('/watch?v=', '')
    url := 'https://www.youtube.com/watch?v=$id'


    {
        println('\e[32m[INFO]\e[m fetch ${url}')
        cmd := 'youtube-dl -f "worstaudio[ext=webm]" -o "%(id)s.%(ext)s" "${url}"'
        println('\e[32m[INFO]\e[m ${cmd}')
        os.system('echo OK')
        os.system(cmd)
        os.system('echo OK2')
    }
    os.system(
        // 'afplay ${id}.m4a &'
        'mplayer -noconsolecontrols -really-quiet ${id}.webm 2>&1 &'
    )
    // pid := C.wait(0)
    // println(pid)

}
