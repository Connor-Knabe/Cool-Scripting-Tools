#!/bin/bash
#Usage: make an alias = ". /mkdirCd.sh"

if [ -z "$1" ];then
	echo "Please enter folder name"
	exit
fi

if [ $1 = "-o" ];then
    mkdir $2
    cd $2
    jcc $2
else
    mkdir $1
    cd $1
fi
