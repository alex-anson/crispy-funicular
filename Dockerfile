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

CMD [ "go", "run", "src/main.go" ]
# CMD [ "ls", "-lah" ]

# docker build . -t my-go-app && docker run -p 3000:10000 my-go-app
# docker-compose up  ...  docker-compose down