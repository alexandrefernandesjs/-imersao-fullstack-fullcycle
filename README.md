# Imersão Full Stack && Full Cycle
 Desafios e aulas desenvolvidos durante a imersão fullstack && fullcycle.

 Participe: https://imersao.fullcycle.com.br

# Microsserviço CodePix
Esse microsserviço tem o objetivo de ser um hub de transações entre os bancos que simularemos durante o projeto.

Também estarão incluidos no projeto alguns desafios que forem requisitados durante a imersão, além do conteúdo da aula.

O projeto original e o conteúdo estritamente das aulas pode ser encontrado aqui: https://github.com/codeedu/imersao-fullstack-fullcycle

# Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

* Faça o clone do projeto
* Tendo o docker instalado em sua máquina apenas execute: `docker-compose up -d`

### Como executar a aplicação
* Acesse o container da aplicação executando: `docker exec -it codepix_app bash`
* Rode `go run cmd/codepix/main.go`

**Importante:** Esse código está sendo disponibilizado conforme o andamento das aulas, logo, é possível que partes do código, principalmente o
main.go, ainda não estejam criados. Caso queira ver o código das aulas mais atualizado, por favor, visite o endereço do projeto original.

### Serviços utilizados ao executar o docker-compose

* Aplicação principal
* Postgres
* PgAdmin
* Apache Kafka
* Criador dos tópicos a serem utilizados pelo Kafka
* Confluent control center
* ZooKeeper
