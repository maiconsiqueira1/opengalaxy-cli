target = /usr/local/bin/og-cli

# Build related variables
buildTime  = `date --iso-8601=ns`
buildHash  = `git log -1 --pretty=format:"%h"`
buildVsn   = $(git tag --points-at $(buildHash))
buildState = `(git diff --quiet && echo 'clean') || echo 'dirty'`

buildFlagsPrefix = github.com/thiagohdeplima/opengalaxy-cli/cmd

build:
	go build -o og-cli -ldflags \
		"-X $(buildFlagsPrefix).buildVsn=$(buildVsn) -X $(buildFlagsPrefix).buildTime=$(buildTime) -X $(buildFlagsPrefix).buildHash=$(buildHash) -X $(buildFlagsPrefix).buildState=$(buildState)" 

install: build
	(test -f $(target) && mv $(target) $(target)-`date +"%s"`) || true
	mv og-cli $(target)