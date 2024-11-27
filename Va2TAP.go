package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Task representa uma tarefa no sistema.
type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"` // pendente, em progresso, concluída
}

// Banco de dados global
var db *gorm.DB

// Inicializa o banco de dados
func initDatabase() {
	var err error

	// Configura o banco de dados usando modernc/sqlite
	db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migração dos modelos
	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("Falha ao migrar modelos:", err)
	}
	log.Println("Banco de dados inicializado com sucesso.")
}

// Criar tarefa
func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível criar a tarefa"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Listar tarefas
func getTasks(c *gin.Context) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível listar as tarefas"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Atualizar tarefa
func updateTask(c *gin.Context) {
	var task Task
	id := c.Param("id")

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível atualizar a tarefa"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Deletar tarefa
func deleteTask(c *gin.Context) {
	var task Task
	id := c.Param("id")

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível deletar a tarefa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso"})
}

// Rota de boas-vindas com HTML
func welcomePage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
		<!DOCTYPE html>
		<html lang="pt-br">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Bem-vindo à API de Tarefas</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					color: #333;
					margin: 0;
					padding: 0;
				}
				.container {
					width: 100%;
					max-width: 800px;
					margin: 0 auto;
					padding: 20px;
					background-color: #fff;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					border-radius: 8px;
				}
				h1 {
					color: #2c3e50;
					text-align: center;
				}
				p {
					font-size: 16px;
					line-height: 1.6;
				}
				ul {
					list-style-type: none;
					padding: 0;
				}
				li {
					font-size: 18px;
					padding: 8px;
					background-color: #ecf0f1;
					margin-bottom: 5px;
					border-radius: 4px;
				}
				footer {
					text-align: center;
					padding: 20px;
					margin-top: 30px;
					background-color: #34495e;
					color: white;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Bem-vindo à API de Gestão de Tarefas!</h1>
				<p>Esta API permite que você gerencie suas tarefas diárias de forma simples e eficiente. Use os seguintes endpoints:</p>
				<ul>
					<li><strong>POST /tasks</strong> - Criar nova tarefa</li>
					<li><strong>GET /tasks</strong> - Listar todas as tarefas</li>
					<li><strong>PUT /tasks/:id</strong> - Atualizar uma tarefa existente</li>
					<li><strong>DELETE /tasks/:id</strong> - Excluir uma tarefa</li>
				</ul>
			</div>
			<footer>
				<p>&copy; 2024 API de Tarefas</p>
			</footer>
		</body>
		</html>
	`)
}

// Roteador principal
func main() {
	// Inicializar o banco de dados
	initDatabase()

	// Configurar rotas
	r := gin.Default()

	// Adicionar a rota de boas-vindas
	r.GET("/", welcomePage)

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	// Iniciar o servidor
	log.Println("Servidor rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor:", err)
	}
}
