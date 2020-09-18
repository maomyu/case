#!/bin/bash
for i in $(seq 1 50)
do
ffmpeg -re -i /Users/mowuya/Desktop/ffmpeg-test/testlong.mp4  -vcodec h264 -acodec aac -f flv "rtmp://58.254.141.76:1966/ly-live/xuqiangv$i?domain=pili-rtmp-publish.tech-idiot.tech" &
#echo $i &
sleep 1;
done
