# Development Container Setup

This document explains how to set up and use the development container (devcontainer) for the Terraform Provider for IBM DataPower. The devcontainer provides a consistent, reproducible development environment using Docker Compose, making it easier for contributors to get started without worrying about local setup variations. It includes a Go development environment and a local instance of IBM DataPower for testing.

## Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) (VS Code) installed.
- The [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) for VS Code.
- [Docker](https://www.docker.com/) installed and running on your machine (with Docker Compose support), or [Podman](https://podman.io/).
- Git installed for cloning the repository.

No additional installations are required inside the container, as everything is handled by the Docker images.

## Getting Started

1. **Clone the Repository**:
   ```
   git clone https://github.com/ScottW514/terraform-provider-datapower.git
   cd terraform-provider-datapower
   ```

2. **Open in VS Code**:
   Open the repository folder in VS Code.

3. **Reopen in Container**:
   - Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on macOS) to open the Command Palette.
   - Type and select **Dev Containers: Reopen in Container**.
   - VS Code will build the Docker images (if not already built) and start the container. This may take a few minutes the first time.

Once the container is running, your workspace will be mounted inside the `go-dev` service at `/workspaces/terraform-provider-datapower/`. You can now develop, build, test, and debug the provider as usual.

## Devcontainer Configuration Files

The devcontainer is defined by the following files in the `.devcontainer/` directory:

### `devcontainer.json`
This file configures the VS Code devcontainer environment:

```json
{
  "name": "Datapower Terraform Provider Dev Container",
  "dockerComposeFile": "docker-compose.yml",
  "service": "go-dev",
  "workspaceFolder": "/workspaces/terraform-provider-datapower/"
}
```

- It specifies the Docker Compose file to use.
- The primary service is `go-dev` (Go development environment).
- The workspace folder is set to the mounted repository path.

### `docker-compose.yml`
This file defines the services used in the devcontainer:

```yaml
services:
  datapower-dev:
    hostname: datapower-dev
    build: datapower
    tty: true
    ports:
      - "5554:5554"
      - "9090:9090"
  go-dev:
    build: go
    env_file: .env
    volumes:
      - ..:/workspaces/terraform-provider-datapower:cached
    command: sleep infinity
```

- **datapower-dev**: Builds a development instance of IBM DataPower from the `datapower/` subdirectory (which contains a `Dockerfile`). It exposes ports 5554 (SOAP) and 9090 (Web GUI).
- **go-dev**: Builds the Go development environment from the `go/` subdirectory (which contains a `Dockerfile`). It loads environment variables from `.env`, mounts the workspace, and runs indefinitely to keep the container alive.

Note: The Dockerfiles for these services are located in `.devcontainer/datapower/` and `.devcontainer/go/`. If you need to customize the builds (e.g., add dependencies), modify these files and rebuild the container.

## Environment Variables (.env File)

The `.env` file provides configuration for the Go environment, such as connection details for the DataPower instance. This file is ignored by Git (via `.gitignore`) to avoid committing sensitive information.

- Copy the provided example to create your own `.env` file:
  ```
  cp .devcontainer/.env.example .env
  ```

- The `.env.example` file looks like this:

  ```shell
  DP_HOSTNAME=datapower-dev
  DP_USERNAME=admin
  DP_PASSWORD=admin
  DP_INSECURE=true
  NO_PROXY=ghcr.io,github.com,datapower-dev
  no_proxy=ghcr.io,github.com,datapower-dev
  DP_ACC_ALL=1
  TF_ACC=1
  ```

- **Key Variables**:
  - `DP_HOSTNAME`, `DP_USERNAME`, `DP_PASSWORD`: Credentials for connecting to the local DataPower instance (defaults to `admin/admin`).
  - `DP_INSECURE`: Set to `true` to allow insecure connections (e.g., self-signed certificates).
  - `NO_PROXY` / `no_proxy`: Exclude certain hosts from proxying (useful in corporate environments).
  - `DP_ACC_ALL`: Enables all acceptance tests for the DataPower Terraform Provider.
  - `TF_ACC`: Set to `1` to run Terraform acceptance tests (which require the real DataPower instance running in the devcontainer).

These variables are required for:
- Running acceptance tests (integration tests that interact with the DataPower instance).
- Debugging in VS Code (e.g., the code expects to connect to DataPower via these vars).
- Unit tests may also rely on these if they simulate interactions, but acceptance tests definitely do.

If you're not running tests or debugging, you can omit the `.env` file, but tests will fail without it.

## Handling Corporate Proxies and Certificates

If you're behind a corporate proxy that requires custom CA certificates:
- Place your CA certificate files (e.g., `.crt` or `.pem`) in the `.devcontainer/go/certs/` directory.
- This directory is ignored by Git (except for the `.gitkeep` file to preserve the folder structure).
- During the `go-dev` image build, these certs will be added to the trust store, allowing Go to handle proxied connections (e.g., to GitHub or package registries).

Rebuild the container after adding certs:
- In VS Code Command Palette: **Dev Containers: Rebuild Container**.

## Accessing the DataPower Instance

The devcontainer includes a running instance of IBM DataPower (`datapower-dev` service):
- **Web GUI**: Access it from your host machine at `https://localhost:9090` (accept any self-signed certificate warnings).
- **Default Credentials**: Username `admin`, Password `admin`.
- **Preconfigured Domain**: The instance includes a domain named `acceptance_test`, which is used by the provider's acceptance tests.

From inside the container (e.g., in the VS Code terminal), you can connect to DataPower using the hostname `datapower-dev`.

## Running Tests and Debugging

- **Unit/Acceptance Tests**: Run with `make test` (may require `.env` vars if they simulate DataPower interactions).
- **Debugging**: Use VS Code's built-in Go debugger. Launch configurations can be added to `.vscode/launch.json`. The debugger will use the `.env` vars to connect to DataPower.

If tests fail due to connection issues:
- Ensure the `datapower-dev` service is running (`docker compose ps`).
- Verify `.env` vars are correct.
- Check for port conflicts on your host (e.g., 9090 or 5554).

## Troubleshooting

- **Build Failures**: If Docker builds fail, check for network issues, proxy settings, or missing dependencies in the Dockerfiles. Run `docker compose build` manually for debugging.
- **Container Not Starting**: Ensure Docker has enough resources (CPU/memory). Check VS Code's Output panel (Dev Containers tab) for logs.
- **Rebuilding**: If you modify devcontainer files, use **Dev Containers: Rebuild Container** from the Command Palette.
- **Exiting the Containers**: The containers runs indefinitely; stop them via Docker or by closing the VS Code window.
- **Performance**: If the container feels slow, allocate more resources to Docker in its settings.

## Contributing

With this setup, you can easily contribute to the provider. Make changes, run tests, and submit pull requests. If you encounter issues with the devcontainer, open an issue in the repository.

If you have custom requirements (e.g., additional tools), feel free to extend the Dockerfiles or propose changes via PR.
