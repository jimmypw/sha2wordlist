BUILDVER := $(shell git describe --tags)

sha2wordlist:
	go build -o sha2wordlist cmd/sha2wordlist/main.go

.PHONY: package
package: sha2wordlist
	@echo Packaging version $(BUILDVER)
	fpm -s dir -t rpm -n sha2wordlist -v $(BUILDVER) --prefix /usr/bin sha2wordlist
