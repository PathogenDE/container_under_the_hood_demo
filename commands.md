
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


 # Do a "manual" container with unshare
sudo unshare --pid --fork --mount-proc ./main


 ## Nsenter into our unshare
 pgrep main
 sudo nsenter -t <PID> -p -m -u /bin/bash


## kill
sudo kill <id>



# Docker doesn't use user namespaces by default

docker run --name=usernsdemo ubuntu sleep infinity
docker inspect usernsdemo | grep Pid
ps aux | grep <123> # look at the user!
docker inspect usernsdemo | grep Pid
cat /proc/<123>/uid_map   # `0` <- container user `0` <- host user `124821512521` <- range

Look into `https://docs.docker.com/engine/security/userns-remap/`

