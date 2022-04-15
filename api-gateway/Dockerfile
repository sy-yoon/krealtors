FROM scratch

# Create and change to the app directory.
#WORKDIR /app/krealtors

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
#COPY ./go.mode ./
#COPY ./go.sum ./
#RUN go mod download

# Copy local code to the container image.
#COPY ./api-gateway/ ./
#COPY ./gms/ ./
#COPY ./google/ ./
#COPY ./protos/ ./
#COPY ./realty/ ./

# Build the binary.
#WORKDIR /app/api-gateway
#RUN go build -v -o server


WORKDIR /app

COPY ./build/gateway ./

CMD ["/app/gateway"]



#git clone https://github.com/sy-yoon/krealtors.git
#go build -v -o build/gateway
#docker build -f api-gateway/Dockerfile -t gateway .
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