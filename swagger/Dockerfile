# Use the official Swagger UI image as base
FROM swaggerapi/swagger-ui:latest

# Set environment variables
ENV SWAGGER_JSON=/app/regolo-api.json
ENV LAYOUT=BaseLayout

# Create directory for API file mounting
RUN mkdir -p /app

# Copy favicon files to the correct directory
COPY regolo-api.json /app/regolo-api.json
COPY favicon-32x32.png /usr/share/nginx/html/favicon-32x32.png
COPY favicon-16x16.png /usr/share/nginx/html/favicon-16x16.png

# Expose port 8080
EXPOSE 8080

# The startup command is already included in the base swagger-ui image
