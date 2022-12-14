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
> #### Run test:
>    >go test -v
> #### Install:
>    >First install the getest module locally using the command "go install github.com/cweill/gotests/...". This will create $GOPATH/bin/gotests.
> #### Usage:
>    >$ gotests [options] PATH ...
> #### Example:
>    >gotests -all ../service/UserService.go > UserService_test.go

> #### Available options:
>    >**-all**                    generate tests for all functions and methods    
>    >
>    >**-excl**                   regexp. generate tests for functions and methods that don't
                              match. Takes precedence over -only, -exported, and -all   
>    >
>    >**-exported**               generate tests for exported functions and methods. Takes
                              precedence over -only and -all   
>    >
>    >**-i**                      print test inputs in error messages   
>    >
>    >**-only**                   regexp. generate tests for functions and methods that match only.
                              Takes precedence over -all   
>    >
>    >**-nosubtests**             disable subtest generation when >= Go 1.7   
>    >
>    >**-parallel**               enable parallel subtest generation when >= Go 1.7.   
>    >
>    >**-w**                      write output to (test) files instead of stdout   
>    >
>    >**-template_dir**           Path to a directory containing custom test code templates. Takes
                              precedence over -template. This can also be set via environment
                              variable GOTESTS_TEMPLATE_DIR
>    >
>    >**-template**               Specify custom test code templates, e.g. testify. This can also
                              be set via environment variable GOTESTS_TEMPLATE
>    >
>    >**-template_params_file**   read external parameters to template by json with file
>    >
>    >**-template_params**        read external parameters to template by json with stdin

## Database engine
Postgresql.

# Run server.
 - go run .../estj/src/main.go

# Build server.
 - go build .../estj/src/main.go
 - This will create an executable file.

# What to do next
~~1. ????????? ?????? ??????.~~
2. ??????, ??????, ???????????? validation ??????.
3. ????????????.
4. ?????????.
5. ????????? ?????? ????????? ?????? ??????. (JWT or API Key or etc)
6. ???????????? ??????. (?????? ????????????, ??????????????? ?????? ???????????? ????????? ?????? ????????????. ????????? ????????? ??????.)
7. ?????? ??????.
8. Web server ???????????? WAS??? ???????????? ??????.
9. Infrastructure
