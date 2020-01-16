FROM golang:latest

COPY main hello-world-app
COPY templates templates

EXPOSE 9090
# Command to run the executable
CMD ["./hello-world-app"]
