#!/bin/bash

for (( i=0; i<30; ++i)); do
    #git tag -d 
    git push origin :v0.0.${i}-alpha
    sleep 1
done