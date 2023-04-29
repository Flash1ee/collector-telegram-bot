FROM golang:1.20

WORKDIR /app

COPY . .

EXPOSE 8082 5432

RUN make build

RUN chmod +x ./app.out

CMD ./app.out -t "TOKEN" -c "./config/config.toml"