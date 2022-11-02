command make image

 - docker build -t pos -f Dockerfile.multistage .

command run image

- docker run -e MYSQL_USER=root -e MYSQL_PASSWORD= -e MYSQL_HOST=docker.for.mac.host.internal -e MYSQL_DBNAME="interviewtest" -p 3030:3030 pos