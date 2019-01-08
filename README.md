# docker compose for web
自分用のサンプルdocker-compose設定です。

# build and up with docker-compose
```
# github.comに設定されている
# SSH_KEY pathはあなたに環境に合わせてください。
# private repositoryを参照する場合に必要。
export SSH_KEY=$(cat {private sshkey path})
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```

# App build and run with docker
```
cd build
docker build --rm -f Dockerfile --build-arg SSH_KEY="$(cat {private sshkey path})" -t app:latest . 
docker run --rm -d --name app app:latest
```
ex)
```
cd build
docker build --rm -f Dockerfile --build-arg SSH_KEY="$(cat {~/.ssh/id_rsa})" -t app:latest . 
docker run --rm -d --name app app:latest
# if you want to direct access this container...
docker run --rm -d --env-file=../config/docker/app/local.env -p 8080:8080 --name app app:latest
```

# Proxy build and run with docker
```
cd build
docker build --rm -f Proxy.Dockerfile -t proxy:latest .
docker run --rm -d --env-file=../config/docker/proxy/local.env -p 80:80 -p 443:443 --link app:app --name proxy proxy:latest
```
