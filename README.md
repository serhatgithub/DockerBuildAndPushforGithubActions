# Docker CI/CD Pipeline

This project automates the building and deployment of the Docker image using GitHub Actions.

## Workflow

For every `push` to the `main` branch, the following steps occur sequentially:

1.  **Test:** The code is tested in different environments (Matrix Strategy).
2.  **Build:** The image is built using the Dockerfile.
3.  **Push:** The image is uploaded to Docker Hub with the `commit SHA` tag.

## Configuration

The following must be defined in the GitHub Repository settings (Secrets and variables):

-   `DOCKERHUB_USERNAME` (Variable): Docker Hub username.
-   `DOCKERHUB_TOKEN` (Secret): Access Token with Read/Write permissions.
