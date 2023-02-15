#!/bin/bash
Project=moon
SrcDir="/tmp/$Project"
WorkDir="/opt/$Project"
CMD="test"
function build() {
  echo 'build'
  cd $SrcDir
  go build -tags=jsoniter "cmd/$CMD.go"
}
function backup() {
  mv "$WorkDir/$CMD" "$WorkDir/$CMD.$(date '+%Y%m%d-%H%M%S')"
}
function copyFresh() {
  echo 'copy fresh.'
  #copy fresh
  cp $SrcDir/$CMD $WorkDir/
  cp $SrcDir/*.yaml $WorkDir/
export opt_env=ga
}

function restart() {
  echo "supervisorctl restart $CMD"
  supervisorctl restart $CMD
}
function validate() {
  sleep 3
  #print check command.
  echo 'curl 127.0.0.1:80/ping'
  curl 127.0.0.1:80/ping
  echo ""
  echo "complete."
}

#main flow
echo '==== start to deploy moon server.'
read -p "select cmd to deploy [
  1:af-push
  2:max-push
]>>" -r mode
echo "mode [$mode] is selected."

case $mode in
1) CMD="af-push" ;;
2) CMD="max-push" ;;
*) echo "UNKNOWN cmd [$mode].Please check and try again." ;;
esac

#build at tmp dir
build
#backup work dir
backup
#copy fresh from tmp to work dir
copyFresh
restart
validate
