FROM go as builder-stage
WORKDIR /root
ADD . .
RUN go test -v ./...
RUN CGO_ENABLED=0 go build -o /root/main .

FROM alpine
COPY --from=builder-stage /root/main /root/main
EXPOSE 8080
CMD [ "/root/main" ]