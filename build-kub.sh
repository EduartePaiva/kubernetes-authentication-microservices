#!/bin/zsh

set -e  # Exit on error
set -u  # Treat unset variables as an error
set -o pipefail  # Exit on first failure in a pipeline

echo "Starting build process..."

# Step 1: build docker images

docker build -t eduartepaiva/kub-dep-auth -f auth.dockerfile .
docker build -t eduartepaiva/kub-dep-users -f users.dockerfile .


# Step 2: Push them to docker hub

docker push eduartepaiva/kub-dep-auth
docker push eduartepaiva/kub-dep-users



echo "Build complete!"                         