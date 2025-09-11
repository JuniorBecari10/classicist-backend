# Classicist

The next-level classical music app.

## How to build

### Dependencies
- Go;
- Bun (handles Vue and TS).

Run:
```sh
make build
```
This will generate an executable file that runs your app.

## For development

Ensure you have installed:

Name|Command to install it
---|---
`go-sqlite3`|`go get github.com/mattn/go-sqlite3`
`reflex`|Arch: `pacman -S reflex`, pip: `pip install reflex`

Run:
```sh
make watch
```
This will reload the program every time you change your source code.

## Image requirements

All images must be of high quality and with aspect ratio of 1:1 (square). Try to find an image that looks good when cut like a banner.
The banner will cut the image at its top.
