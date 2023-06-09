package delivery

import (
	"fmt"
	"simple-payment/config"
	"simple-payment/delivery/controller"
	"simple-payment/delivery/middleware"
	"simple-payment/manager"
	"simple-payment/util/authenticator"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine         *gin.Engine
	host           string
	useCaseManager manager.UseCaseManager
	tokenService   authenticator.AccessToken
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("/api")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewUserController(publicRoute, s.useCaseManager.UserUseCase())
	controller.NewCustomerController(publicRoute, s.useCaseManager.CustomerUseCase(), tokenMdw)
}

func (s *Server) Run() {
	s.initController()

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
	tokenService := authenticator.NewAccessToken(conf)

	rm := manager.NewRepositoryManager(infra)
	ucm := manager.NewUseCaseManager(rm, tokenService)

	return &Server{
		engine:         router,
		host:           host,
		useCaseManager: ucm,
		tokenService:   tokenService,
	}
}
