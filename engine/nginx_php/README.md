## Start up memcached 
`docker run --name my-memcached -d memcached`

## Build php and run
`docker build -t my-phpfpm:7.0 ./php`

`docker run --name my-phpfpm --link my-memcached:memcached -d my-phpfpm:7.0`



## Build nginx and run
Now that phpfpm and memcached are running and linked together let's run a nginx proxy in front of the phpfpm instance.
You can build an ubuntu base with nginx installed like so:

`docker build -t my-ubuntu-nginx:1.4 ./ubuntu`

`docker run -it -p 80:80 --link my-phpfpm:phpfpm my-ubuntu-nginx:1.4`

Or you can use the official nginx image as a base:

`docker build -t my-nginx:1.9 ./nginx`

`docker run -it -p 80:80 --link my-phpfpm:phpfpm my-nginx:1.9`
