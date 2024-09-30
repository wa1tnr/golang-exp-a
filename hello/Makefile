# Mon 30 Sep 18:15:32 UTC 2024

build:	hello.go
	go build

# go build -race

run:	hello.go
	@echo '    ' go run - only source files remain\; nothing compiled remains
	go run .

clean:
	@rm -rf ./hello

hello:	build

# fark: build + exec <from a file>
fark:	hello
	@./hello 

exec:	fark

# errors:
# https://www.gnu.org/software/make/manual/html_node/Errors.html

# cruft:
# @echo '    ' use \'make build\' not \'make run\'

# end.
