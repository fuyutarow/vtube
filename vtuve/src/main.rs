use std::os::unix::io::{AsRawFd, FromRawFd};
use std::process::{Child, Command, Stdio};
use scraper::Html;

let selector = Selector::parse("li").unwrap();

for element in fragment.select(&selector) {
    assert_eq!("li", element.value().name());
}

fn main() {
    let url = "https://www.youtube.com/watch?v=VBpXOOPK6-E";
    // let mut proc = Command::new("youtube-dl")
    // let mut proc = Command::new(format!(
    let mut proc = Command::new("youtube-dl")
        .arg("-f 249")
        .arg("-o '%(id)s.%(ext)s'")
        .arg(url)
        .spawn()
        .expect("failed to execute process");
    proc.wait();
}
