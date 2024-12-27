# Use the official OpenResty base image
FROM openresty/openresty:alpine

# Copy the Nginx configuration file
COPY nginx.conf /etc/openresty/nginx.conf

# Expose port 8080 for OpenResty
EXPOSE 8080

# Run OpenResty in the foreground
CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]