#!/bin/bash

# docker buildx create --name builder --driver docker-container --use
docker buildx build --platform linux/amd64,linux/arm64 -t zooneon/echo-server:latest --push .