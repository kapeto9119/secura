FROM node:16-alpine

WORKDIR /app

# Create a placeholder for now
RUN echo "console.log('Frontend placeholder');" > index.js

# Use a simple server to respond to requests
RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-p", "3000", "."] 