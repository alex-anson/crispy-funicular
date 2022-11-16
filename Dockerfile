FROM golang:1.19-buster

# Adds an entry to `etc/passwd` inside the image's filesystem space, among other things.
# "alex" will own all installed files.
RUN useradd -ms /bin/bash alex

# Working directory INSIDE the container.
WORKDIR /home/alex/code

# Transfers files from outside to inside container.
COPY --chown=alex:golang . .

CMD [ "go", "run", "src/main.go" ]

# docker build . -t my-go-app && docker run -p 3000:10000 my-go-app