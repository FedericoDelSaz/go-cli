# top_numbers CLI Tool

This repository provides the **top_numbers** CLI tool that reads an input file containing numbers, retrieves the N largest numbers, sorts them in descending order, and outputs the result.

## top_numbers Docker Image

The **top_numbers** Docker Image can be built and pushed to Docker Hub using GitHub Actions. This section explains how to trigger the Docker image creation workflow, build the Docker image, and push it to your specified Docker repository.

### Triggering the Docker Image Creation Workflow

You can manually trigger the **top_numbers Docker Image** workflow to build and push the Docker image to your Docker repository.

#### Steps to Trigger the Workflow

1. Go to the **Actions** tab of this GitHub repository.
2. Select the **top_numbers Docker Image** workflow from the left-hand sidebar.
3. Click on the **Run Workflow** button on the right side.
4. Enter the following inputs when prompted:
    - **imageTag**: Example `latest` or a specific version (e.g., `3.14.3`).
    - **repositoryName**: The Docker repository where the image should be pushed (e.g., `docker.io/felodel/top_numbers`).

#### Workflow Steps:
- **Checkout repository**: The repository is checked out to ensure the latest code is used for building the Docker image.
- **Login to Docker Hub**: Authenticates using your Docker Hub credentials. Ensure you have the `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets set in your GitHub repository.
- **Build Docker image**: Builds the Docker image with the specified tag (e.g., `latest` or version number).
- **Push Docker image**: Pushes the Docker image to the specified Docker repository.

### Prerequisites
Before running the workflow, make sure you have the following set up:

- A **Docker Hub** account.
- **DOCKERHUB_USERNAME** and **DOCKERHUB_TOKEN** set as secrets in your GitHub repository:
    - **DOCKERHUB_USERNAME**: Your Docker Hub username.
    - **DOCKERHUB_TOKEN**: A personal access token (use Docker Hub’s token instead of your password for security reasons).

### Example Workflow

1. **Trigger Workflow**: Click the **Run Workflow** button on the GitHub Actions tab.
2. **Provide Inputs**:
    - `imageTag`: `0.0.4`
    - `repositoryName`: `docker.io/felodel/top_numbers`
3. The workflow will build the Docker image and push it to `felodel/top_numbers:0.0.4`.

### Docker Image Usage

Once the image is built and pushed to Docker Hub, you can pull and use the `top_numbers` Docker image as follows:

```bash
docker pull <your-docker-repository>:<image-tag>
