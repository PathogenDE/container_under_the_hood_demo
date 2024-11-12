
 # Explore process in /proc/

 ls -l /proc/1234/exe
 cat /proc/1234/cmdline
 cat status

 pstree -p


# running container and exploring it

docker run --name pidteller_container -e SECRET=env1r0mentS3cretrun! pidteller

## Finding Pid of the container

docker inspect pidteller_container | grep Pid

## Finding secrets (env)

docker inspect pidteller_container | grep SECRET


## Finding the fs (overlay2)

docker inspect pidteller_container | grep MergedDir



# Finding clone3 using strace 

strace -f -e clone3 docker run --name clone3finder alpine echo "I started" && sleep 5 && echo " -- I finished --"


 # unshare
sudo unshare --pid --fork --mount-proc ./main


 # Nsenter
 pgrep main
 sudo nsenter -t <PID> -p -m -u /bin/bash


# kill
kill <id>


