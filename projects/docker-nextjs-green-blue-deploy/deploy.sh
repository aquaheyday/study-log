#!/bin/bash

# Variables
VHOST_CONF_PATH="/etc/httpd/conf.d/httpd-vhost.conf"
SSL_CONF_PATH="/etc/httpd/conf.d/ssl.conf"
BLUE_CONTAINER="nextjs-app-blue"
GREEN_CONTAINER="nextjs-app-green"
BLUE_PORT=6960
GREEN_PORT=6961

# Determine active container
if grep -q "ProxyPass / http://127.0.0.1:$BLUE_PORT/" $VHOST_CONF_PATH; then
    echo "Blue is active. Switching to Green..."
    ACTIVE_PORT=$BLUE_PORT
    TARGET_PORT=$GREEN_PORT
    ACTIVE_CONTAINER=$BLUE_CONTAINER
    TARGET_CONTAINER=$GREEN_CONTAINER
else
    echo "Green is active. Switching to Blue..."
    ACTIVE_PORT=$GREEN_PORT
    TARGET_PORT=$BLUE_PORT
    ACTIVE_CONTAINER=$GREEN_CONTAINER
    TARGET_CONTAINER=$BLUE_CONTAINER
fi

# Build and deploy the target container
echo "Building and deploying $TARGET_CONTAINER..."
docker-compose build $TARGET_CONTAINER
docker-compose up -d $TARGET_CONTAINER

# Wait for the new container to be healthy
echo "WaiUMUMting for $TARGET_CONTAINER to become healthy..."
sleep 10  # Adjust this based on app startup time

# Update VirtualHost configuration
echo "Updating VirtualHost configuration to point to $TARGET_PORT..."
sudo sed -i "s|http://127.0.0.1:$ACTIVE_PORT/|http://127.0.0.1:$TARGET_PORT/|g" $VHOST_CONF_PATH

# Update VirtualHost configuration
echo "Updating VirtualHost configuration in $VHOST_CONF_PATH and $SSL_CONF_PATH..."
sudo sed -i "s|http://127.0.0.1:$ACTIVE_PORT/|http://127.0.0.1:$TARGET_PORT/|g" $VHOST_CONF_PATH
sudo sed -i "s|http://127.0.0.1:$ACTIVE_PORT/|http://127.0.0.1:$TARGET_PORT/|g" $SSL_CONF_PATH


# Reload Apache to apply the changes
echo "Reloading Apache..."
sudo apachectl graceful

# Stop and remove the old container
echo "Stopping and removing $ACTIVE_CONTAINER..."
docker-compose stop $ACTIVE_CONTAINER
docker-compose rm -f $ACTIVE_CONTAINER

echo "Deployment completed. Active service: $TARGET_CONTAINER"
