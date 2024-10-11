# Wed  2 Oct 12:19:21 UTC 2024
# was:
#   Tue  1 Oct 12:40:17 UTC 2024
#   Mon 30 Sep 18:15:32 UTC 2024

build:	runes.go extra.go
	@echo make build ran
	go build

initONLYONCE:
	@go mod init example/runes

# go build -race

tidy:
	@echo make tidy ran
	@go mod tidy

# @echo '    ' go run - only source files remain\; nothing compiled remains
run:	runes.go extra.go
	go run .

# test is a synonym for 'run' in this Makefile
test:	run

# note: should tidy be called earlier or later than clean
clean:	tidy
	@rm -rf ./runes

distclean:	clean

# exec is at OS shell level
exec:	build
	@./runes

# no reason or motive to 'install' this program anywhere anytime for any reason - established. ;)
install:
	@echo '    ' nothing to do - no interest in \'installing\' this program.

# errors:
# https://www.gnu.org/software/make/manual/html_node/Errors.html

# cruft:
# @echo '    ' use \'make build\' not \'make run\'

# doco - see https://pandoc.org/demos.html  for proper use syntax ;)
doco:	notes.md
	pandoc -s notes.md 2>/dev/null | lynx -stdin

# given as: 'godoc builtin'  in Miek Gieben's book
# today in debian:
# prerequisite:  'apt install golang-golang-x-tools'

# found $PAGER undefined!

PAGER=/usr/bin/less
godocc:
	go doc builtin | ${PAGER}

mangodoc:
	lynx https://manpages.debian.org/bookworm/golang-golang-x-tools/godoc.1.en.html

# end.
