# Metrics

Service to collect information from node exporter and cadvisor about servers for my blog (https://frkam.dev)

## Local development

Copy and fill .env file

`cp .env.example .env && nano .env`

Run project

go run .

## Deployment

Copy and fill .env file in `./deployment`

`cp .env.example ./deployment/.env && cd deployment && nano .env`

Then run docker compose

`sudo docker compose up -d`
