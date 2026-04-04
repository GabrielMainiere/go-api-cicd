# 🚀 API em Go

## 📌 Sobre o Projeto
API simples desenvolvida em GO como forma de treinamento para criação do dockerfile, compose e ci/cd

Pré-requisitos
- Go instalado (>= 1.20)

Passo a passo
- cd api-go
- go mod tidy
- go run main.go

A API estará disponível em:
- http://localhost:3000

## 🐳 (Atividade 1) Rodando com Docker

Você deverá criar um Dockerfile para essa aplicação

### Requisitos:

- Utilizar imagem oficial do Go
- Build da aplicação
- Usar multi-stage build
- Expor porta 3000
- Executar a API


## 🐳 (Atividade 2) Rodando com Compose

Você deverá modificar a aplicação para fazer acesso ao banco de dados. Crie um docker compose para executar o PostreSQL, o PGAdmin e a aplicação em GO atraves do dockerfile que você criou

### Requisitos:

- Utilizar o dockerfile criado na atividade 1
- Criar um docker compose com:
  - A aplicação em go
  - O banco PostgreSQL
  - O PGAdmin




## ⚙️ (DESAFIO) CI/CD

Crie um CI/CD no github actions com as seguintes etapas

- CI (Integração Contínua)
  - Build da aplicação
  - Testes unitários
  - Testes de integração
  - Lint
  - Análise de qualidade de código (SonarQube)
  - SAST (Semgrep ou Checkmarx ou Fortify, etc)

- Container
  - Docker Lint
  - Build da imagem
  - Scan de vulnerabilidades (Trivy)
  - Push da imagem no dockerhub

- CD (Entrega Contínua)
   - Deploy em homologação com Render
   - DAST (OWASP ZAP)
   - Criação da aprovação manual
   - Deploy em produção