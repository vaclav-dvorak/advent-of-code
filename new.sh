#! /bin/sh
#? Usage: ./new.sh 2022 7
set -e
source .env
day=$2
if [[ ${#day} -lt 2 ]] ; then
    day="00${day}"
    day="${day: -2}"
fi
DIR=$1/day${day}
if [ -d ${DIR} ]; then
 echo "${DIR} directory exists.";
 exit 1;
fi
cp -R template ${DIR}
curl --cookie "session=${ADVENT_OF_CODE_COOKIE}" https://adventofcode.com/${1}/day/${2}/input -o ${DIR}/input.txt
