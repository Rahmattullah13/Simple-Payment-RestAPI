package delivery

import (
	"fmt"
	"simple-payment/config"
	"simple-payment/manager"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine         *gin.Engine
	host           string
	useCaseManager manager.UseCaseManager
}

func (s *Server) Run() {
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running on port : ", s.host)
	}
}

func NewServer(host string) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	conf := config.NewConfig()
	infra := manager.NewInfraManager(conf)

	rm := manager.NewRepositoryManager(infra)
	ucm := manager.NewUseCaseManager(rm)

	return &Server{
		engine:         router,
		host:           host,
		useCaseManager: ucm,
	}
}
