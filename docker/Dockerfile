# Use Go as the base image
FROM golang:1.23

# Install Node.js 20 and npm
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Install Reflex
 RUN go install github.com/cespare/reflex@v0.3.1

# Copy Reflex configuration and startup script
COPY reflex.conf /
COPY start.sh /

# Set the entrypoint to Reflex
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]