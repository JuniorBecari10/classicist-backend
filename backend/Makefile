# Dependencies:
# Go: pacman -S go
# Reflex: go install github.com/cespare/reflex@latest

FILE=../backend

build:
	go build -o $(FILE) main.go

run:
	go run main.go -- yes

watch:
	reflex -s -r '\.go$$' make run
