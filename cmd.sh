#!/bin/sh

init() {
  [ -z $1 ] && usage
  no=$1
  mkdir -p "problem${no}" && 
  echo "package problem${no}" > "problem${no}/sol.go" && 
  echo "package problem${no}" > "problem${no}/sol_test.go"
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
