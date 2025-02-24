# top_numbers CLI and Docker Image

This project provides a CLI tool `top_numbers` that reads an input file, retrieves the N largest numbers, and outputs them in descending order. It also provides a Docker image for easy usage and deployment.

## CI/CD Pipeline

The repository includes two key GitHub Actions workflows:

- **Build-Test**: This workflow builds and tests the Go project and optionally generates the `top_numbers` CLI tool.
- **docker-image.yml**: This workflow builds and pushes the Docker image for the project to a Docker registry.

### 1. Using the **Build-Test** Workflow to Generate the `top_numbers` CLI

The **Build-Test** workflow can be triggered manually through GitHub Actions to build the `top_numbers` CLI and run tests.

#### Triggering the Workflow:
1. Go to the **Actions** tab of the GitHub repository.
2. Click on the **Build-Test** workflow on the left side.
3. Click on the **Run Workflow** button on the right side.
4. You will be prompted to enter an input for `generateCLI`:
    - **generateCLI**: Set this to `true` if you want to generate the `top_numbers` CLI tool. Set it to `false` to just build the project.
    - Default: `false`

#### Steps:
- If `generateCLI` is `true`, the workflow will:
    1. Build the project located in the `sre-cli-tool` directory.
    2. Upload the generated `top_numbers` binary as an artifact.

- If `generateCLI` is `false`, the workflow will:
    1. Skip generating the CLI and instead build and test the `sre-cli-tool`.

#### Example Output:
- After the workflow completes, if `generateCLI` is `true`, you can download the `top_numbers` binary artifact from the workflow page.

### 2. Using the **docker-image.yml** Workflow to Build and Push the Docker Image

The **docker-image.yml** workflow helps you build and push the Docker image for this project.

#### Triggering the Workflow:
1. Go to the **Actions** tab of the GitHub repository.
2. Click on the **docker-image.yml** workflow on the left side.
3. Click on the **Run Workflow** button on the right side.
4. You will be prompted to enter the following inputs:
    - **imageTag**: Specify the tag for the Docker image (e.g., `latest` or `v1.0.0`).
    - **repositoryName**: Specify the Docker repository name (e.g., `docker.io/yourusername/top_numbers`).

#### Steps:
- The workflow will:
    1. Build the Docker image using the `Dockerfile` present in the repository.
    2. Push the Docker image to the specified Docker registry (e.g., Docker Hub) using the provided repository name and image tag.

#### Example Command to Trigger Workflow:

```bash
# Triggering docker-image.yml workflow manually with inputs
# Image Tag: latest, Repository Name: docker.io/yourusername/top_numbers
