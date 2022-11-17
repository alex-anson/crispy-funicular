FROM golang:1.19-buster

# Adds an entry to `etc/passwd` inside the image's filesystem space, among other things.
# "alex" will own all installed files - IF that's the user you transfer ownership to within the COPY instruction.
RUN useradd -ms /bin/bash alex

# Working directory INSIDE the container.
WORKDIR /home/alex/code

# Transfers files from outside to inside container.
# Change file ownership to 'alex' user.
# COPY --chown=<user>:<group> <hostPath> <containerPath>
COPY --chown=alex:alex . .

# Created .dockerignore & added go.mod & go.sum to it. Docker won't copy over either 
# of these files now. (...and I deleted them from my local machine)
# Create the go.mod file
RUN go mod init main
# Add necessary dependencies
RUN go get github.com/gorilla/mux@v1.8.0
# Make sure go.mod matches the source code in the module.
RUN go mod tidy

CMD [ "go", "run", "src/main.go" ]
# CMD [ "ls", "-lah" ]

# docker build . -t my-go-app && docker run -p 3000:10000 my-go-app
# docker-compose up  ...  docker-compose down