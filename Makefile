
build:
	go build -o bin/vtube main.go

install:
	cp bin/vtube /usr/local/bin

uninstall:
	rm /usr/local/bin/vtube
