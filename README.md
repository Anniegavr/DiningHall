# DiningHall
https://hub.docker.com/r/anniegavr/dininghall - docker container here

Resources for learning and coding: https://zetcode.com/golang/getpostrequest/
https://youtu.be/W5b64DXeP0o - Go API code, part 1 
https://youtu.be/YMQUQ6XQgz8 - Go API code, part 2

This container is for the dining hall service; it sends http requests to the RESTaurant server.
The two docker containers communicate with each other over a network bridge, being assigned a network IP address.
![image](https://user-images.githubusercontent.com/56108881/134954813-21dfe11f-8630-4eb1-b472-74d8dfee1020.png)

Example of requests being sent back and forth:
![image](https://user-images.githubusercontent.com/56108881/134955815-1518d10e-50b2-4e29-9dae-3ed2d2617fdd.png)
