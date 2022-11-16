FROM golang:1.19-buster

# Not ideal. 0 means root. if this were meant for prod, and there was an exploit, they'd be root (... no idea what i'm talking about)
# Tried to use 1212, didn't know how to get around the permission denied error - failed to initialize build cache at /.cache/go-build: mkdir /.cache: permission denied
USER 0

# Working directory INSIDE the container.
WORKDIR /home/go/code

# Transfers files from outside to inside container.
# Eeeeekk root.
COPY --chown=0:golang . .

CMD [ "go", "run", "src/main.go" ]

# docker build . -t my-go-app && docker run -p 3000:10000 my-go-app