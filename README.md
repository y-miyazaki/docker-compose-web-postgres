# docker compose and docker for web 
This repository is sample docker for golang and postgres and redis and fluentd and s3.
I created 2 way build.

* docker-compose
* docker 

# docker-compose

## build and up
```
# .env.local is local sample.
# if you want to use docker-compose, please copy .env.local to .env.
# .env file use for docker-compose.yml.
cp -p .env.local .env
# github.comに設定されている
# SSH_KEY pathはあなたに環境に合わせてください。
# private repositoryを参照する場合に必要。
export SSH_KEY=$(cat {private sshkey path})
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```

# docker
## Proxy and App add netowork with docker
```
docker network create app_proxy
```

## App build and run with docker
```
docker build --rm -f build/Dockerfile --build-arg SSH_KEY="$(cat {private sshkey path})" -t app:latest . 
docker run --rm -d --net app_proxy --name app app:latest
```
ex)
```
docker build --rm -f build/Dockerfile --build-arg SSH_KEY="$(cat ~/.ssh/id_rsa)" -t app:latest . 
docker run --rm -d --net app_proxy --env-file=config/docker/app/local.env --name app app:latest
```

### If you want to direct access this container.
```
docker run --rm -d --net app_proxy --env-file=config/docker/app/local.env -p 8080:8080 --name app app:latest
```

## Proxy build and run with docker
```
docker build --rm -f build/Proxy.Dockerfile -t proxy:latest .
docker run --rm -d --net app_proxy --env-file=config/docker/proxy/local.env -p 80:80 -p 443:443 --link app:app --name proxy proxy:latest
```

# Access web
```
curl localhost
```

# Etc...

```
# remove all docker images
alias dockerallrmi='docker rmi -f $(docker images | awk '\''{print$3}'\'')'
# stop all docker containers
alias dockerallps='docker stop $(docker ps -a -q) | docker rm $(docker ps -a -q)'
# remove all volumes
alias dockervolrm='docker volume rm $(docker volume ls -qf dangling=true)'
```