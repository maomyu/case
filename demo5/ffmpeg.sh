#!/bin/bash
for i in $(seq 1 100)
do
ffmpeg -re  -i ~/Desktop/jiangziya_hls.mp4  -codec copy  -f flv "rtmp://10.200.20.28:1966/jztest4/test$i?domain=pili-publish.jztest4.cloudvdn.com" &
#echo $i &
sleep 1;
done