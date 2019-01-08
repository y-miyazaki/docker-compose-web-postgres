FROM golang:1.11.4-alpine3.8

ARG SSH_KEY

WORKDIR /go/src/github.com/y-miyazaki/docker-compose-web-postgres

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash gcc musl-dev git mercurial openssh curl && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

# Fresh for rebuild on code change, no need for production
# RUN go get -u github.com/pilu/fresh && \
RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u bitbucket.org/liamstask/goose/cmd/goose

# if your application have private library with github.com.
RUN mkdir -p /root/.ssh && \
    echo "$SSH_KEY" > /root/.ssh/id_rsa && \
    chmod 0600 /root/.ssh/id_rsa && \
    eval `ssh-agent` && \
    ssh-add /root/.ssh/id_rsa && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts && \
    echo -e "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig && \
    echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

COPY . ./

RUN dep ensure -v -vendor-only=true && \
    go build

FROM golang:1.11.4-alpine3.8

COPY --from=0 /go/src/github.com/y-miyazaki/docker-compose-web-postgres /go/src/github.com/y-miyazaki/docker-compose-web-postgres

RUN apk update && \
    apk upgrade && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

CMD ["./docker-compose-web-postgres"]
