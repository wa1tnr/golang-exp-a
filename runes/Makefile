# Tue  1 Oct 12:40:17 UTC 2024
# was:
#   Mon 30 Sep 18:15:32 UTC 2024

build:	hello.go
	go build

initONLYONCE:
	@go mod init example/runes

# go build -race

# @echo '    ' go run - only source files remain\; nothing compiled remains
run:	hello.go
	go run .

# test is a synonym for 'run' in this Makefile
test:	run

clean:
	@rm -rf ./hello

distclean:	clean

# exec is at OS shell level
exec:	build
	@./hello 

# no reason or motive to 'install' this program anywhere anytime for any reason - established. ;)
install:
	@echo '    ' nothing to do - no interest in \'installing\' this program.

# errors:
# https://www.gnu.org/software/make/manual/html_node/Errors.html

# cruft:
# @echo '    ' use \'make build\' not \'make run\'

# end.
