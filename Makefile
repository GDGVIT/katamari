windows:
	@echo "Building for windows"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/katamari.exe
linux:
	@echo "Building for linux"
	go build -o ./bin/linux/katamari
	@ cd bin/linux && ./katamari
all:
	@echo "Building for every OS and Platform"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/katamari.exe
	GOOS=linux GOARCH=386 go build -o ./bin/linux/katamari
	GOOS=freebsd GOARCH=386 go build -o ./bin/freebsd/katamari-bsd
	GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/katamari-mac
run:
	go run .
global:
	go install .
release:
	gh release create $v './bin/windows/katamari.exe' './bin/linux/katamari' './bin/mac/katamari-mac'