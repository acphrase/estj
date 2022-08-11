# Introduction
ESTJ Project.

# Getting Started
This is a Gin in golang application.

## Package layout
1. *.dataaccesslayer package contains all the entities and repositories organized into structs.
2. *.service package contains all the business logic "service" objects. These service objects will utilize the repositories and entities.
3. *.api package contains all the controllers and implements the API protocol. These controllers will use service objects and handle data transformation for the API input/output.

## Logging
1. Use zap made by Uber.
2. Among logging frameworks, I was thinking of Zerolog and Zap, but I chose Zap because golang users use Zap a lot.

## Testing
TODO

## Database engine
Postgresql.

# Run server.
 - go run .../estj/src/main.go

# Build server.
 - go build .../estj/src/main.go
 - This will create an executable file.

# What to do next
1. 테스트 코드 관련.
2. 회원가입.
3. 로그인.
4. 로그인 유저 확인을 위한 방법. (JWT or API Key or etc)
5. 미들웨어 구성. (유저 확인이나, 단순코드나 이런 여러가지 확인을 위한 미들웨어. 우선은 로그인 유저.)
6. 도커 구성.
7. Web server 사용하여 WAS로 연결하는 방법.
8. Infrastructure
