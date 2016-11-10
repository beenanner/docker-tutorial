## Build nodejs and run
`docker build -t my-nodejs:latest ./nodejs`

`docker run --name my-nodejs -d my-nodejs:latest`



## Build nginx and run
Now that nodejs is running let's run a nginx proxy in front of the nodejs instance.

`docker build -t my-nginx:1.9 ./nginx`

`docker run -it -p 80:80 --link my-nodejs:nodejs my-nginx:1.9`
