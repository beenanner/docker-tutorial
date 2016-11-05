## Build golang and run
`cd golang && docker build -t my-golang:latest .`

`docker run --name my-golang -d my-golang:latest`



## Build nginx and run
Now that golang is running let's run a nginx proxy in front of the golang instance.

`cd nginx && docker build -t my-nginx:1.9 .`

`docker run -it -p 80:80 --link my-golang:golang my-nginx:1.9`
