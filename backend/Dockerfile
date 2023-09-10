FROM golang:1.20

WORKDIR /usr/src/app
COPY backend/go.mod backend/go.sum ./
RUN go mod download && go mod verify
COPY ./backend .

CMD ["go", "run", "./src/server.go"]
