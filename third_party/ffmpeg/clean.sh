#!/bin/bash

set -e
set -x

echo "Doing clean..."

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
pushd "$DIR/remote"

make clean

popd
