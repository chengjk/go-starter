#!/bin/bash

Project=moon
# work dir at target server.
WorkDir=/tmp/$Project
# generate version by time
echo 'generate version'
VERSION=$(date "+%Y%m%d-%H%M%S")
sed -i .tmp "s/version:.*/version: $VERSION/g" ./conf*.yaml

# define functions

## env-hk,海外香港
function executeAg() {
  sed -i .tmp 's/.*opt_env.*/export opt_env=ga/g' ./deploy.sh
  HOSTS=(
    'ec2-user@16.163.177.227'
  )
  uploadServers "$HOSTS" '~/.ssh/aws-ag.pem'
}

#upload to server one by one
function uploadServers() {
  HOSTS=$1
  PEM=$2
  # shellcheck disable=SC2068
  for h in ${HOSTS[@]}; do
    echo ">>>> $h upload  start."
    rsync -av -e "ssh -i $PEM" --exclude={'.git/','.idea/','logs/','*.tmp'} ../$Project/ "$h":$WorkDir
    echo "<<<< $h upload  complete."
  done
}

#main flow
echo '==== upload to servers'
read -p "input env [
1:海外-hk
]>>" -r mode
echo ""
case $mode in
1) executeAg ;;
*) echo ":UNKNOWN env.Please check and try again." ;;
esac
rm -f *.tmp
echo '==== upload all server complete.'
