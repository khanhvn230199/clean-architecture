package master

import (
	"fmt"
	"github.com/example-golang-projects/clean-architecture/cmd/server/master/config"
	"github.com/example-golang-projects/clean-architecture/services/master/dependency"
	"log"
	"net/http"
)

type MasterService struct {
}

func (s *MasterService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.InitMasterRouter().ServeHTTP(w, r)
}

func NewMasterService(cfg config.Config) *MasterService {
	dependency.InitMasterDependency(cfg)
	return &MasterService{}
}

func RunMasterService(cfg config.Config) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServicePort),
		Handler: NewMasterService(cfg),
	}
	err := server.ListenAndServe()
	switch err {
	case nil, http.ErrServerClosed:
	default:
		log.Fatal(err, nil, nil)
	}
}
