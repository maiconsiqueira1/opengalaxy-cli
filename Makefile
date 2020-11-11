target = /usr/local/bin/og-cli

build:
	go build -o og-cli

install: build
	(test -f $(target) && mv $(target) $(target)-`date +"%s"`) || true
	mv og-cli $(target)