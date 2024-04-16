# Start from the official golang image based on Debian
FROM golang:1

# Set the working directory inside the container
WORKDIR /app

# Install necessary packages, including SQLite with FTS5 support
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    sqlite3 \
    gcc \
    libc-dev \
    libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev libnss3 libxss1 libasound2 libxtst6 xauth xvfb \
    && rm -rf /var/lib/apt/lists/*

# Enable CGO
ENV CGO_ENABLED=1

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN DEV=PROD go build --tags="fts5" -o web ./main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web"]
