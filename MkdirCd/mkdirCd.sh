#!/bin/bash
#Usage: make an alias = ". /mkdirCd.sh"

openFlag=false;
echo "Test"
echo $openFlag

while getopts ":o:" option; do

	case "$option" in
    o) 
        echo "TRUE"
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

echo "Dir Name: " + $2

echo $openFlag

if [ "$openFlag" = true ] ; then 
    echo 'Is true'
    echo $openFlag
    mkdir $2
    cd $2
    jcc $2

else 
    echo 'Is false'
    mkdir $1
    cd $1
fi
