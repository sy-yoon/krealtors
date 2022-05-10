FROM golang:1.17.8-alpine3.15 as builder

# Create and change to the app directory.
WORKDIR /app/krealtors

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY ./go.mod ./
RUN go mod download
#RUN go mod tidy -v

# Copy local code to the container image.
COPY ./gms-region ./gms-region
COPY ./gms ./gms
COPY ./google ./google
COPY ./protos ./protos


# Build the binary.
#RUN go mod tidy
RUN go mod tidy -compat=1.17
RUN go build -v -o ./build/region ./gms-region

FROM alpine:3.9
COPY --from=builder /app/krealtors/gms-region/server.env /app/krealtors/server.env
COPY --from=builder /app/krealtors/build/region /app/krealtors/region
CMD ["/app/krealtors/region"]


# 이미지 생성 명령어 (현 파일과 같은 디렉토리에서)
# docker build -t {이미지명} .

# 컨테이너 생성 & 실행 명령어
# docker run --name {컨테이너명} -v $(pwd):/home/node/app -p 80:9000 {이미지명}

#git clone https://github.com/sy-yoon/krealtors.git
#go build -v -o build/gateway
#docker build -f api-gateway/Dockerfile -t gateway .
#docker stop {컨테이너명}
#docker container rm -f {컨테이너명}
#docker images rm -f {image}
#docker run --nmae {containerName} -i -t {imageName}


#docker tag gateway:latest coolnice/gateway:latest
#docker login (coolnice/1qaz2wsx!!)
#docker push coolnice/gateway:latest






# 이미지 생성 명령어 (현 파일과 같은 디렉토리에서)
# docker build -t {이미지명} .

# 컨테이너 생성 & 실행 명령어
# docker run --name {컨테이너명} -v $(pwd):/home/node/app -p 80:9000 {이미지명}



# CMD 에서 작업
# SET GOOS=linux
# SET GOARCH=amd64
# go buid -v -o ../build/server