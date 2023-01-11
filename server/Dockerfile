FROM golang:1.19-alpine3.16

RUN apk add --update nodejs npm
RUN echo $(node --version)
RUN echo $(npm --version)

RUN npm i -g redoc-cli

RUN mkdir -p /app/esoft

COPY server /app/esoft/server

WORKDIR /app/esoft/server

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -o docs --ot json -q -g main.go

RUN redoc-cli build \
    		-o docs/index.html docs/swagger.json \
    		--options.sortOperationsAlphabetically=true \
    		--options.sortTagsAlphabetically=true

RUN go build -o server main.go

EXPOSE 4000

CMD ["./server"]