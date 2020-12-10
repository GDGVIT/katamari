windows:
	@echo "Building for windows"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/katamari.exe
linux:
	@echo "Building for linux"
	go build -o ./bin/linux/katamari
	cd bin/linux && ./katamari
all:
	@echo "Building for every OS and Platform"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/katamari.exe
	GOOS=linux GOARCH=386 go build -o ./bin/linux/katamari
	GOOS=freebsd GOARCH=386 go build -o ./bin/freebsd/katamari
	GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/katamari
	@echo "Zipping for release"
	@tar -czvf bin/releases/katamari_linux.tar.gz bin/linux/katamari LICENSE 
	@tar -czvf bin/releases/katamari_win.tar.gz bin/windows/katamari.exe LICENSE 
	@tar -czvf bin/releases/katamari_mac_amd64.tar.gz bin/mac/katamari LICENSE 
	@tar -czvf bin/releases/katamari_bsd.tar.gz bin/freebsd/katamari LICENSE 
run:
	go run .
global:
	go install .

# do not use
release:
	gh release create $v ${find /bin/releases -type f -printf "%f "}