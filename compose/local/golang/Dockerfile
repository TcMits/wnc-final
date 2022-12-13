FROM golang:1.19-bullseye

COPY ./compose/local/golang/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

COPY ./compose/local/golang/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start


WORKDIR /wnc-final

COPY go.mod go.sum /wnc-final/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["/entrypoint"]
