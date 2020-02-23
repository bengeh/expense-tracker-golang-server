#!/bin/bash
imageName=expensetracker-image
containerName=expensetracker-container

docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep $imageName )

docker build -t $imageName -f Dockerfile  .

echo Delete old container...
docker rm -f $containerName

echo Run new container...
docker run -d -p 41960:8000 --name $containerName $imageName