
 # Explore process in /proc/

 ls -l /proc/1234/exe
 cat /proc/1234/cmdline
 cat status

 pstree -p


 # unshare
sudo unshare --pid --fork --mount-proc ./main


 # Nsenter
 pgrep main
 sudo nsenter -t <PID> -p -m -u /bin/bash


# kill
pgrep -f unshare
kill <id>


# Docker

docker run --name nsenter nsenter_demo

/var/lib/docker/overlay2

 docker inspect nsenter | grep MergedDir

 .../merged


 touch iwashere

 
