package handlers

import (
	toDoService "effectiveMobile/internal/fio/service"
	"net/http"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		id, err := s.toDoService.Create(r.Context(), toDoService.CreateParams{
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{ID: id}, http.StatusOK)
	}
}
