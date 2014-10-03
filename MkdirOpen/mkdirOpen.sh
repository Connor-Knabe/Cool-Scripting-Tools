#!/bin/bash
if [ -z "$1" ];then
	echo "Please enter folder name"
	exit
fi
mkdir $1
cd $1