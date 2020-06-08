#!/bin/bash
./bitcoin send --from paradise --to wang --m 10 
./bitcoin send --from wang --to yao --m 5.3
./bitcoin send --from yao --to wang --m 3.1
./bitcoin getBit --ad wang
./bitcoin getBit --ad yao
./bitcoin getBit --ad paradise
echo "Hope wang:7.8 yao: 2.2 paradise: 26"
