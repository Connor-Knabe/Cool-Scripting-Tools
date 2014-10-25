#!/bin/bash
#Usage: make an alias = ". /mkdirCd.sh"

openFlag=false;

while getopts :o: option
do
	case "$option" in
    o) openFlag = true
		


if [ -z "$1" ];then
	echo "Please enter folder name"
	exit
fi
mkdir $1
cd $1

if $openFlag = true ; then
	jcc $1
fi 