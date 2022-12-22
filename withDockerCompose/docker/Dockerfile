FROM debian:buster AS compiler

ARG GO_VERSION
RUN apt update && apt install -y curl
WORKDIR /tmp
RUN URL=https://storage.googleapis.com/golang/go1.${GO_VERSION}.linux-amd64.tar.gz \
    && curl ${URL} -o go.tar.gz  \
    && tar -zxf go.tar.gz  \
    && rm -rf go.tar.gz  \
    && mv go /go

ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin
# If you enable this, then gcc is needed to debug your app
ENV CGO_ENABLED 0

FROM compiler AS dev

WORKDIR /app
EXPOSE 8000 40000
RUN apt update && apt install -y curl vim fish sudo git \
    && groupadd -g 1000 docker-user \
    && useradd -d /home/docker-user -s /bin/bash -u 1000 -g 1000 docker-user \
    && usermod -aG sudo docker-user && echo "docker-user:1234" | sudo chpasswd \
    && mkdir /home/docker-user \
    && chown -R docker-user:docker-user /home/docker-user \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && ls -s /go/bin/dlv /usr/local/bin \
    && chown -R docker-user:docker-user /go

FROM compiler AS compile

WORKDIR /app
COPY ./app .
RUN go build -ldflags "-s -w" -o final.sh

FROM debian:buster AS prod

WORKDIR /
EXPOSE 8000
COPY --from=compile /app/final.sh /
CMD ["./final.sh"]