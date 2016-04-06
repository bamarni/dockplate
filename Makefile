.PHONY: clean build release

clean:
	rm -rf build

build: clean
	for os in "darwin" "linux" "windows" ; do \
		mkdir -p build/$$os; \
		GOOS=$$os GOARCH=amd64 go build -o build/$$os/dockplate ./dockplate.go; \
	done

release: build
	for os in "darwin" "linux" "windows" ; do \
	mkdir -p build/releases; \
	tar -cvzf build/releases/dockplate-$$os-x64.tar.gz -C build/$$os dockplate; \
	done

