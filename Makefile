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
	@tar -czf bin/releases/katamari_linux.tar.gz LICENSE -C bin/linux katamari
	@tar -czf bin/releases/katamari_win.tar.gz LICENSE -C  bin/windows katamari.exe
	@tar -czf bin/releases/katamari_mac_amd64.tar.gz LICENSE -C bin/mac katamari 
	@tar -czf bin/releases/katamari_bsd.tar.gz LICENSE -C bin/freebsd katamari 
run:
	go run .
global:
	go install .

# do not use
release:
	gh release create $v 'bin/releases/katamari_linux.tar.gz' 'bin/releases/katamari_win.tar.gz' 'bin/releases/katamari_bsd.tar.gz' 'bin/releases/katamari_mac_amd64.tar.gz' 

gdg:
	gh release create $v 'bin/releases/katamari_linux.tar.gz' 'bin/releases/katamari_win.tar.gz' 'bin/releases/katamari_bsd.tar.gz' 'bin/releases/katamari_mac_amd64.tar.gz' -R GDGVIT/katamari