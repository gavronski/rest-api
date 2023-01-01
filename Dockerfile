FROM golang:1.18

# Set the working directory to the /var/www/html directory
WORKDIR /var/www/html

# Copy go.mod go.sum into the image
COPY go.mod /var/www/html
COPY go.sum /var/www/html

# Install modules inside the image
RUN go mod download

# Copy the source code into the image
COPY . /var/www/html

# Install soda
RUN go install github.com/gobuffalo/pop/v6/soda@latest

# Compile Go app
RUN go build -o /app ./cmd/web/*.go

# Run the Go app when the container starts
CMD ["/app"]
