

.PHONY: build

build:
	make -f cmd/leo/Makefile build

ctl:
	make -f cmd/leoctl/Makefile build

install:
	make -f cmd/leoctl/Makefile install

reload:
	make -f cmd/leo/Makefile reload