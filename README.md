
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
```

## Implantação com Docker

Certifique-se de ter o Docker instalado em sua máquina antes de prosseguir.

1. Clone este repositório.
2. Configure o arquivo `config.json` conforme mencionado anteriormente.
3. Abra um terminal na raiz do projeto.
4. Execute o comando `docker-compose up -d` para iniciar o contêiner em segundo plano.

## Verificando a Porta 32781

Se desejar acessar a API de fora da sua máquina local, verifique se a porta 32781 está aberta e acessível para o mundo externo. Você pode fazer isso usando um serviço online ou um utilitário local como o `nmap`.

Por exemplo, no terminal, você pode executar o comando:

```bash
nmap -p 32781 localhost
```

Isso verificará se a porta 32781 está aberta em `localhost`. Certifique-se de substituir `localhost` pelo endereço IP da máquina onde a API está sendo executada, se aplicável.
































# Consumindo a API de Envio de E-mails

Este guia fornece instruções sobre como consumir a API de Envio de E-mails usando PHP e Python.

## Endpoints Disponíveis

- **POST** `/send-email`: Envia um e-mail usando os parâmetros fornecidos.

## Consumindo a API com PHP

### Exemplo de Uso

Aqui está um exemplo simples de como enviar um e-mail usando a API em PHP:

```php
<?php
$url = 'http://localhost:32781/send-email';
$data = array(
    'emissor' => 'seu_email@gmail.com',
    'assunto' => 'Assunto do E-mail',
    'destinatario' => 'destinatario@example.com',
    'mensagem' => 'Corpo da mensagem do e-mail',
    'usuario' => 'Seu Nome'
);

$options = array(
    'http' => array(
        'header' => "Content-type: application/json\r\n",
        'method' => 'POST',
        'content' => json_encode($data)
    )
);

$context = stream_context_create($options);
$result = file_get_contents($url, false, $context);

if ($result === FALSE) {
    echo "Erro ao enviar e-mail";
} else {
    echo "E-mail enviado com sucesso!";
}
?>
