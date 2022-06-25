#!/usr/bin/env bash
set -ex

newFolder() {
	dirname=$(date +%Y-%m-%d)
	if [ ! -d "$dirname" ]; then
		mkdir -p $dirname
	fi
}

copyFiles() {
	srcdir=kata-src
	destdir=$(date +%Y-%m-%d)
	cp -r $srcdir/* $destdir
}

watch() {
	destdir=$(date +%Y-%m-%d)
	nodemon -e go -w ./$destdir -x "cd ${destdir} && go test -v ./... || true"
}

startDay() {
	newFolder
	copyFiles
	watch
}

startDay
