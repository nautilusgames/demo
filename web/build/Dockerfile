# Use the official Nginx image as a base image
FROM nginx:alpine

# Copy the static website files to the Nginx web directory
COPY . /usr/share/nginx/html

# Expose port 8080
EXPOSE 8080

# Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]