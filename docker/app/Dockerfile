FROM golang:1.22.3

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
# RUN go install github.com/go-sql-driver/mysql

CMD ["air"]

# CMD ["go", "run", "main.go"]
