#! /bin/sh

prog_dir=$(cd "$(dirname "$0")"; pwd)
prog_arch=$(getconf LONG_BIT)
export LD_LIBRARY_PATH="${prog_dir}/lib"

case "$1" in
    stop)
        exec ./watcher $1
        ;;
    info)
        exec ./watcher $1
        ;;
     status)
        exec ./watcher $1
        ;;
     restart)
        exec ./watcher $1
        ;;
    *)
        exec ./watcher start
        ;;
esac