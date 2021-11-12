#
# Build UI project
#
FROM node:16-alpine as ui-builder

WORKDIR /usr/src/ui

COPY ui/ .

RUN yarn

RUN yarn build --modern

#
# Build revel app
#
FROM golang:1.17-bullseye as app-builder

WORKDIR /usr/src/r_res

RUN go get -u github.com/revel/revel
RUN go install github.com/revel/cmd/revel@latest

COPY go.mod .
COPY go.sum .

COPY app ./app

COPY conf ./conf
RUN sed -i 's/mongo.host =.*/mongo.host = mongo-db/g' conf/app.conf
RUN sed -i 's/http.addr =.*/http.addr = 0.0.0.0/g' conf/app.conf

COPY messages ./messages
COPY tests ./tests

COPY --from=ui-builder /usr/src/ui/dist ./public

RUN go mod download

RUN revel build -a . -t build

#
# Deploy
#
FROM debian:bullseye-slim

WORKDIR /opt/r_res

COPY --from=app-builder /usr/src/r_res/build .

ENTRYPOINT ["/opt/r_res/run.sh"]