#!/bin/bash

set -e
set -x

echo "Doing build..."

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
BUILD_DIR="$DIR/build"
pushd "$DIR/remote"

git reset --hard
git clean -f


./configure \
	--prefix="$BUILD_DIR" \
	--pkg-config-flags="--static"  \
	--extra-libs="-lpthread -lm" \
	--bindir="$BUILD_DIR/bin" \
	--enable-debug \
	--disable-stripping \
	--enable-libx264 \
	--enable-gpl \
	 && make -j 4 && make install


git reset --hard
git clean -f

popd
