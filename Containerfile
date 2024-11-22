FROM golang:1.22 AS build
WORKDIR /src
COPY . /src
RUN go build -o /bin/advent .

FROM gcr.io/distroless/base
COPY --from=build /bin/advent /bin/advent