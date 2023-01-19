FROM golang:alpine AS total
WORKDIR /app
COPY . /app/
RUN go build -o main 

FROM alpine
LABEL "name"="ascii-art-web-dockerize"
LABEL maintaner="Azel_Nura"
WORKDIR /app-2
COPY --from=total /app/main /app-2/
COPY --from=total /app/ui/ /app-2/ui
COPY --from=total /app/ascii-art/files/ /app-2/ascii-art/files
EXPOSE 8080
CMD [ "./main" ]