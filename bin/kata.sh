#!/usr/bin/env bash
set -e

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

ensuredeps() {
	if ! nodemon -v >/dev/null; then
		echo "nodemon not found"
		npm install -g nodemon
	fi
}

watch() {
	destdir=$(date +%Y-%m-%d)
	nodemon -e go -w ./$destdir -x "cd ${destdir} && go test -v ./... || true"
}

startDay() {
	newFolder
	copyFiles
	ensuredeps
	watch
}

startDay
