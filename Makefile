# nnn-select
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build

clean:
	rm -f nnn-select

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f nnn-select $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/nnn-select

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/nnn-select

.PHONY: all build clean install uninstall
