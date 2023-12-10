# FROM golang:1.21.3

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the auth server code into the container
# COPY ./auth-server .

# # Build the Go application
# RUN go build -o auth-server ./cmd

# # Expose the port that the auth server will run on
# EXPOSE 8001

# # Command to run the auth server
# CMD ["./auth-server"]


FROM golang:1.21.3

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

#RUN go mod tidy


