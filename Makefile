
build:
	go build -o vtube play.go

install:
	mv vtube /usr/local/bin

uninstall:
	rm /usr/local/bin/vtube

v\:build:
	v build -o vtube play.v


