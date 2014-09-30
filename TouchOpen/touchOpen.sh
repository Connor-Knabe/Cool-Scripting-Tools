#!/bin/bash
if [ -z "$1" ];then
	echo "Please enter file name"
	exit
fi
touch $1
open $1