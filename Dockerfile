FROM docker.io/library/golang:1.18 as builder

# Set up the working directory
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.{mod,sum} files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


#-- Final -- 
FROM scratch

WORKDIR /notapp

ENV SUPER_SECRET=123password
COPY super.secret /notapp/

# Copy the binary from the build stage
COPY --from=builder /app/main /notapp/


# Set the command for the container
CMD ["./main"]
