#!/bin/bash
#Usage: make an alias = ". /mkdirCd.sh"

openFlag=false;

while getopts :o: option
do
	case "$option" in
    o) 
        openFlag = true
        ;;
    *)
        echo "Something went wrong"
        ;;
    esac
		
done

if [ -z "$1" ];then
	echo "Please enter folder name"
	exit
fi


if [ "$openFlag" ] ; then 
    echo 'Is true'
    mkdir $1
    cd $1
else 
    echo 'Is false'
fi


if $openFlag ; then
	jcc $1
fi 