### Docker Volume

`docker volume create --name my-volume`

`docker run -d --name ubuntu -v my-volume:/my-volume ubuntu:trusty sleep 1h`

`docker cp ./my-file.txt ubuntu:/my-volume/my-file.txt`

`docker exec -it ubuntu bash`

`cat /my-volume/my-file.txt`

`exit`

`docker rm -f ubuntu`

`docker run -it --name alpine -v my-volume:/my-volume alpine:latest cat /my-volume/my-file.txt`

Volumes are stored in `/var/lib/docker/volumes`

`cat /var/lib/docker/volumes/my-volume/_data/my-file.txt`

You can also mount existing directories for real-time editing outside of docker

`docker run -it --name ubuntu -v $(pwd)/volumes:/my-volume ubuntu:trusty bash`