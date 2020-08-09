NAME=ffmpeg
#
#echo "---------------"
#
#echo 'killing ->' $NAME
#
#ID=`ps -ef | grep "$NAME" | grep -v "grep" | awk '{print $2}'`
#echo 'found ID list:' $ID
#for id in $ID
#do
## 杀掉进程
#kill $id &
#echo "killed $id" &
#done
#echo "---------------"
ps -ef | grep ffmpeg | grep -v "grep" | awk '{print $2}' | xargs kill -9