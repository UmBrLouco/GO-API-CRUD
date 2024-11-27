# GO-API-CRUD
Projeto de tópicos avançados de programação.
Para executar o projeto, insira os seguintes comandos no terminal da IDE após fazer o download do arquivo .GO:


1- go get modernc.org/sqlite
2- go get gorm.io/driver/sqlite
3- go get modernc.org/sqlite


Comandos para criar, deletar, alterar;

criar:
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\":\"Trabalho de Topicos em programação\",\"description\":\"Fazer uma API para gestão de tarefas\",\"priority\":\"Alta!!\",\"status\":\"pendente\"}"

Listar todas as tarefas:
curl -X GET http://localhost:8080/tasks

Atualizar uma tarefa existente (PUT);
curl -X PUT http://localhost:8080/tasks/ID_DA_TAREFA -H "Content-Type: application/json" -d "{\"title\":\"Tarefa Atualizada\",\"description\":\"Descrição atualizada\",\"priority\":\"Baixa\",\"status\":\"em progresso\"}"

Deletar uma tarefa (DELETE);
curl -X DELETE http://localhost:8080/tasks/ID_DA_TAREFA


