# Use the official Debian-based Go image as the base image
FROM debian:latest

# Set the working directory in the container
WORKDIR /app

# Copy the compiled Go binary into the container
COPY httpmqtt /app/

# Expose any necessary ports
EXPOSE 8080

# Define the command to run your Go program
CMD ["/app/httpmqtt"]
