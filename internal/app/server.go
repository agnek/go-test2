package app

import (
	"github.com/gin-gonic/gin"
)

func StartServer(b *Bank, config Config) error {
	s := NewServer(b)

	r := gin.Default()
	secured := r.Use(NewAuthFunc(config.ApiToken))
	secured.POST("/user", s.CreateUser)
	secured.GET("/user/:id/balance", s.GetBalance)
	secured.POST("/transfer", s.Transfer)

	return r.Run(config.Bind)
}

func NewServer(b *Bank) *Server {
	return &Server{b}
}

type Server struct {
	app *Bank
}

type CreateUserRequest struct {
	Id      string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (s *Server) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	c.BindJSON(&req)
	s.app.CreateUser(req.Id, req.Balance)
	c.String(200, "{}")
}

func (s *Server) GetBalance(c *gin.Context) {
	userId := c.Param("id")
	balance := s.app.GetBalance(userId)
	c.String(200, "{\"balance\": %f, \"user_id\": %s}", balance, userId)
}

type TransferRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"float64"`
}

func (s *Server) Transfer(c *gin.Context) {
	var req TransferRequest
	c.BindJSON(&req)
	err := s.app.Transfer(req.From, req.To, req.Amount)
	if err != nil {
		c.Status(500)
	} else {
		c.String(200, "{}")
	}

}
