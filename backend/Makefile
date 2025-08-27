# Dependencies:
# Go: pacman -S go
# Reflex: go install github.com/cespare/reflex@latest

build:
	go build main.go

run:
	go run main.go

watch:
	reflex -s -r '\.go$$' make run
