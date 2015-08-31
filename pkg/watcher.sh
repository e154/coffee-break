#! /bin/sh

prog_dir=$(cd "$(dirname "$0")"; pwd)
prog_arch=$(getconf LONG_BIT)
export LD_LIBRARY_PATH="${prog_dir}/lib" && ./watcher