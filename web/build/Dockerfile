# Use the official Nginx image as a base image
FROM nginx:alpine

# Copy the static website files to the Nginx web directory
COPY . /usr/share/nginx/html

# Expose port 80 to the outside world
EXPOSE 80

# Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]