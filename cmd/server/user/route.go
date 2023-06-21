package user

import "github.com/go-chi/chi/v5"

func (s *UserService) InitUserRouter() (router chi.Router) {
	router = chi.NewRouter()
	router.Route("/api/user", func(r chi.Router) {
		r.Route("/role", func(r chi.Router) {
			r.Get("/", s.roleHandler.CreateRole)
		})
		r.Route("/permission", func(r chi.Router) {
			r.Get("/", s.roleHandler.CreateRole)
		})
	})
	return router
}
