# Dependencies:
# Go: pacman -S go
# Reflex: go install github.com/cespare/reflex@latest

FILE=../setup_bin

build:
	go build -o $(FILE) main.go

run: build
	trap "rm -fr $(FILE)" EXIT; \
		$(FILE) yes

watch:
	reflex -s -r '\.go$$' make run
