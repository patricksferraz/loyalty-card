# Kubernetes Deployment Instructions

This document provides step-by-step instructions for deploying the Loyalty Card service on Kubernetes.

## Prerequisites
- A running Kubernetes cluster
- `kubectl` configured to communicate with your cluster
- Docker registry credentials (if using a private registry)

## Create Secrets

### Environment Variables Secret
1. Fill in the `.env` file in `k8s/.env` using `k8s/.env.example`.
2. Run the following command to create a secret from the `.env` file:
   ```sh
   kubectl create secret generic loyalty-card-secret --from-env-file k8s/.env
   ```

### Docker Registry Secret
Run the following command to create a secret for accessing your Docker registry:
```sh
kubectl create secret docker-registry regsecret \
--docker-server=$DOCKER_REGISTRY_SERVER \
--docker-username=$DOCKER_USER \
--docker-password=$DOCKER_PASSWORD \
--docker-email=$DOCKER_EMAIL
```
Where:
- `$DOCKER_REGISTRY_SERVER`: URL for the registry
- `$DOCKER_USER`: Registry username
- `$DOCKER_PASSWORD`: Registry password
- `$DOCKER_EMAIL`: Optional, any email

## Deploy the Application
Run the following command to deploy all Kubernetes resources:
```sh
kubectl apply -f ./k8s
```

## Additional Notes
- Ensure your `.env` file is properly configured before creating secrets.
- For more details, refer to the [README.md](../README.md) in the root directory.
