CHANGELOG 	:= gateway
BUILDTIME 	:= $(shell date "+%F %T")
MKPATH		:= $(abspath $(lastword $(MAKEFILE_LIST)))
MKDIR		:= $(shell dirname ${MKPATH})
SOURCE 		:= ${MKDIR}/
DESTDIR 	:= $(abspath ${MKDIR}/../../target/build/bin/)
PROJECTDIR  := $(abspath ${MKDIR}/../../)
TARGET 		:= leo
PREFIX 		:= main
SUB         := $(shell sh ${MKDIR}/../../scripts/version.sh leo)
VERSION 	:= 0.0.1.${SUB}
COMMIT      := "${CHANGELOG} (${SUB})"
LDFLAGS 	:=  "-X '${PREFIX}.Version=${VERSION}'     		\
  -X '${PREFIX}.Built=${BUILDTIME}'							\
  -X '${PREFIX}.ChangeLog=${CHANGELOG}'					\
  -s														\
  -w														\
"

default: build

build:
	mkdir -p ${DESTDIR}
	go build -ldflags ${LDFLAGS} -o ${DESTDIR}/${TARGET} ${SOURCE}

version: build
	${DESTDIR}/${TARGET} version

install: build
	cp ${DESTDIR}/${TARGET} $(HOME)/.peon/${TARGET}

reload: install
	launchctl unload ~/Library/LaunchAgents/peon.leo.plist
	launchctl load ~/Library/LaunchAgents/peon.leo.plist