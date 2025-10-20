CHANGELOG 	:= init
BUILDTIME 	:= $(shell date "+%F %T")
MKPATH		:= $(abspath $(lastword $(MAKEFILE_LIST)))
MKDIR		:= $(shell dirname ${MKPATH})
SOURCE 		:= ${MKDIR}/
DESTDIR 	:= $(abspath ${MKDIR}/../../target/build/bin/)
PROJECTDIR  := $(abspath ${MKDIR}/../../)
TARGET 		:= leoctl
PREFIX 		:= main
SUB         := $(shell sh ${MKDIR}/../../scripts/version.sh leoctl)
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
	cp ${DESTDIR}/${TARGET} "/Users/iuz/Library/Application Support/Alfred/Alfred.alfredpreferences/workflows/user.workflow.FFEC67C5-AB6D-4BB1-BD51-E2CF0D24C615/"