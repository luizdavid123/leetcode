#!/bin/sh

init() {
  [ -z "$1" ] && usage
  title="$1"
  mkdir -p "${title}" && 
  echo "package sol" > "${title}/sol.go" &&
  echo "package sol" > "${title}/sol_test.go"
}

save() {
  [ -z "$1" ] && usage
  msg="$1"
  git add . &&
  git commit -m "$msg" &&
  git push origin master
}

usage() {
  echo "usage: sh $0 [init no] | [save msg]"
  exit 0
}

# no cmd args
[ $# -eq 0 ] && usage

cmd=$1
case $cmd in
  init)
    init $2
    ;;
  save)
    save "$2"
    ;;
  help | *)
    usage
    ;;
esac
