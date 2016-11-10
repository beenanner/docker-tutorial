## Build golang and run
`docker build -t my-golang:latest ./golang`

`docker run --name my-golang -d my-golang:latest`



## Build nginx and run
Now that golang is running let's run a nginx proxy in front of the golang instance.

`docker build -t my-nginx:1.9 ./nginx`

`docker run -it -p 80:80 --link my-golang:golang my-nginx:1.9`
