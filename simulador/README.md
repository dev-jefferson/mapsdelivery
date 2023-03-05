# mapsdelivery

## Simulador de posições

Aplicação desenvolvida em linguagem GO para gerar as posições do mapsdelivery


Executar container:
- docker-compose up -d

Acessar container:
- docker exec -it simulator bash

Executar dentro do container
- go run main.go

Utilizando GO Mod para o controle de gerenciamento de pacotes 
- go mod init github.com/dev-jefferson/mapsdelivery
