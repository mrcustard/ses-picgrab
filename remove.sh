#!/bin/bash

filelist=$(ls -l | grep -E "10922|10932" | awk '{ print $9 }')

for i in $filelist; do 
    mv $i notpic/
    echo "moved: $i"
done
