# golight - userspace driver for managing brightness on nvidia macbooks

include config.mk

SRC = main.go

all: options golight

options:
	@echo "CC	= ${CC}"
	@echo "DESTDIR	= $(DESTDIR)"

golight:
	${CC} -o golight

clean:
	rm -rf golight

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f golight ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/golight
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/VERSION/${VERSION}/g" < golight.1 > ${DESTDIR}${MANPREFIX}/man1/golight.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/golight.1

uninstall:
	rm -rf ${DESTDIR}${PREFIX}/bin/golight
	rm -rf ${DESTDIR}${manprefix}/man1/golight.1

.PHONY: all options clean install uninstall
