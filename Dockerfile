FROM golang:1.19-buster

# Adds an entry to `etc/passwd` inside the image's filesystem space, among other things.
# "alex" will own all installed files - IF that's the user you transfer ownership to within the COPY instruction.
RUN useradd -ms /bin/bash alex

# Working directory INSIDE the container.
WORKDIR /home/alex/code

# Transfers files from outside to inside container.
# Change file ownership to 'alex' user.
# COPY --chown=<user>:<group> <hostPath> <containerPath>
COPY --chown=alex:alex ./go .

# Docker won't copy over any of the files specified in the .dockerignore file.
# Create the go.mod file
RUN go mod init main
# Add necessary dependencies
RUN go get github.com/gorilla/mux@v1.8.0
# Make sure go.mod matches the source code in the module.
RUN go mod tidy

# TODO: make this idea work, or take it out
ENV CONTAINER=true 

CMD [ "go", "run", "main.go" ]
# CMD [ "ls", "-lah" ]

# docker build . -t my-go-app && docker run -p 3000:10000 my-go-app
# docker-compose up  ...  docker-compose down

# BELOW = vs-code auto-generated
# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=goapi Version=0.0.1
# EXPOSE 3000
