#!/bin/bash
#Usage: make an alias = ". /mkdirCd.sh"
if [ -z "$1" ];then
	echo "Please enter folder name"
	exit
fi
mkdir $1
echo "Made"
echo "HII"
dirpath=$1
cd $dirpath