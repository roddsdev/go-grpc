

# Código do DESAFIO 1 Imersão Full Stack && Full Cycle

Não programo em GOLANG, mas aceitei o desafio de fazer este servidor de GRPC simples, simulando um cadastro e consulta de produtos.

## Como usar os Endpoints

Após ter subir o docker (ver abaixo) e subir o servidor GRPC (ver abaixo), você pode:
- Abrir um novo terminal
- Acesse o container da aplicação executando: `docker compose exec app bash`
- Executar o comando para entrar no client de GRPC chamado EVANS: `evans -r repl`
- Executar o comando `package example.grpc` 
- Executar o comando `service ProductService` 
- Executar o comando `call $endpoint` 

Existem 3 $endpoints:
- CreateProduct
- FindProduct (passando uuid)
- FindProducts (lista todos os productos cadastrados na base)

## Como subir o docker

Utilizei Docker, conforme ensinado nas aulas.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como subir o servidor GRPC
- Acesse o container da aplicação executando: `docker compose exec app bash`
- Rode `go run main.go`
- Neste momento servidor estará pronto para aceitar conexões

### Serviços utilizados ao executar o docker-compose

- Aplicação principal (GRPC)
- Postgres
- PgAdmin

