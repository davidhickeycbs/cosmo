#!/bin/bash

# Variables
AWS_ACCOUNT_ID="884875474022"
AWS_REGION="us-east-1"
ECR_REPO_NAME="cosmo-router"
IMAGE_TAG="latest"

# Authenticate Docker to AWS ECR
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# Check if the ECR repository exists, and create it if it doesn't
aws ecr describe-repositories --repository-names $ECR_REPO_NAME --region $AWS_REGION > /dev/null 2>&1

if [ $? -ne 0 ]; then
  echo "ECR repository $ECR_REPO_NAME does not exist. Creating..."
  aws ecr create-repository --repository-name $ECR_REPO_NAME --region $AWS_REGION
  echo "ECR repository $ECR_REPO_NAME created."
else
  echo "ECR repository $ECR_REPO_NAME already exists."
fi

# Build the Docker image
docker build -t $ECR_REPO_NAME -f custom.Dockerfile .

# Tag the Docker image
docker tag $ECR_REPO_NAME:latest $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

# Push the Docker image to ECR
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

echo "Docker image pushed to ECR: $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG"