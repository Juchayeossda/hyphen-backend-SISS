# FROM golang:1.21.1

# ENV INSTALL_PATH=/app

# ENV RUN_PATH=./app

# RUN mkdir -p $INSTALL_PATH

# WORKDIR $INSTALL_PATH

# ADD ./app $INSTALL_PATH

# RUN go get

# RUN go build -ldflags="-s -w" -o app .

# CMD ./app

# FROM golang:1.21.1

# WORKDIR /hyphen-backend-SISS

# COPY . .

# RUN go mod tidy

# RUN go build -ldflags="-s -w" -o /run .

# CMD /run

# 기본 이미지를 가져옵니다. Go 언어를 사용하는 경우 golang 이미지를 사용합니다.
FROM golang:1.21.1

# 작업 디렉토리를 설정합니다.
WORKDIR /app

# 현재 디렉토리의 모든 파일을 컨테이너의 /app 디렉토리로 복사합니다.
COPY . .

# Go 어플리케이션을 컴파일합니다.
RUN go build -o myapp main.go

# 컨테이너에서 실행될 명령을 지정합니다.
CMD ["./myapp"]
