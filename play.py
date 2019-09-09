import subprocess
import time


cmd = 'youtube-dl -f "worstaudio[ext=webm]" -o "%(id)s.%(ext)s" "https://www.youtube.com/watch?v=QASCoM5n1-o"'
proc = subprocess.Popen(
        cmd.split(),
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        encoding='utf8',
    )


cmd = 'mplayer pSBk3QVY3cA.webm'
#  with subprocess.Popen(
#          cmd.split(),
#          #  shell=True,
#          stdin=subprocess.PIPE,
#          stdout=subprocess.PIPE,
#          stderr=subprocess.PIPE,
#          encoding='utf8',
#          ) as proc:

proc = subprocess.Popen(
        cmd.split(),
        #  shell=True,
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        encoding='utf8',
    )

print("#1")
time.sleep(5)
#  proc.stdin.write('p\n'.encode())
proc.stdin.write('p')

print("#2")
time.sleep(5)
#  proc.stdin.write('p')
#  proc.stdin.write(b'p')

print("#3")
time.sleep(10)

print("#4")
