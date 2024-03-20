# API de Envio de E-mails

Este é um servidor de API simples para enviar e-mails usando SMTP. Aqui está uma breve visão geral dos arquivos e funcionalidades relevantes:

- `main.go`: O arquivo principal contendo o código-fonte da API.
- `Dockerfile`: Arquivo para a construção de uma imagem Docker para implantação.
- `docker-compose.yml`: Arquivo de composição do Docker para orquestração de contêineres.
- `config.json`: Arquivo de configuração que contém informações confidenciais como endereço de e-mail e senha para autenticação SMTP.

## Como Configurar

Antes de implantar a API, você precisa configurar o arquivo `config.json` com as credenciais corretas. Aqui está um exemplo de como deve ser o formato:

```json
{
  "email": "seu_email@gmail.com",
  "password": "sua_senha_de_email"
}
