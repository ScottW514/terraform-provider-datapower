FROM mcr.microsoft.com/devcontainers/go:1-1.24-bookworm

COPY certs/* /usr/local/share/ca-certificates/
RUN update-ca-certificates

RUN wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
RUN apt update && apt install terraform -y

COPY files/.terraformrc /home/vscode/
