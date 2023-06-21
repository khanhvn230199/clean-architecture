package master

import "github.com/go-chi/chi/v5"

func (s *MasterService) InitMasterRouter() (router chi.Router) {
	router = chi.NewRouter()
	return router
}
