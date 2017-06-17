#!/bin/bash

#constants
MAX_ITERATIONS=10
FILE_NAME="output.txt"

#script
touch $FILE_NAME #guarantees that file exists
rm $FILE_NAME #so it doesnt print unnecessary info

#runs the program the desired number of times
for i in `seq 1 $MAX_ITERATIONS`;
do
	start=$( date +"%T" )
	go run etapa1.go
	end=$( date +"%T" )
	echo "-------------------------ITERATION #"$i"-------------------------" >> $FILE_NAME
	echo "Start time: "$start >> $FILE_NAME
	echo "Finished time: "$end >> $FILE_NAME
	echo "Duration: "$(($(date -d $end +%s) - $(date -d $start +%s))) "second(s)" >> $FILE_NAME
	echo"" >> $FILE_NAME
done