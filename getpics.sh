#!/bin/bash
#set -x
for i in $(cat NewPartsList.txt) ; do 
  baseurl="https://www.temnantco.com/content/dam/tennant/tennantco/products/parts%20and%20supplies/$i.jpg"
  echo "Getting: $i"
  curl -s -XGET --limit-rate 200k -O $baseurl
done
