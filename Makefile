
build:
	go build -o bin/vtube src/play.go

install:
	cp bin/vtube /usr/local/bin

uninstall:
	rm /usr/local/bin/vtube

v\:build:
	v build -o vtube play.v


