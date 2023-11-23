# Build image
FROM node:16-alpine as vue-build
WORKDIR /app

# Install dependencies
COPY package.json yarn.lock ./
RUN yarn

# Build app
COPY ./ .
RUN yarn build

# Final image
FROM nginx:alpine
EXPOSE 80

# Copy the nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy the build from the previous stage
COPY --from=vue-build /app/dist /usr/share/nginx/html
