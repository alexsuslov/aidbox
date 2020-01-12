All:
	go build -ldflags "-X main.version=`git describe --abbrev=0`" -o aidbox
