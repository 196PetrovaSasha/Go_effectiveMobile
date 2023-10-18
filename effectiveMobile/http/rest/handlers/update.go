package handlers

import (
	toDoService "effectiveMobile/internal/fio/service"
	"errors"
	"github.com/fir1/rest-api/pkg/erru"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name        *string `json:"name"`
		Surname     *string `json:"surname"`
		Patronymic  *string `json:"patronymic"`
		Nationality *string `json:"nationality"`
		Age         *int    `json:"age"`
		Gender      *string `json:"gender"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		req := request{}
		err = s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.toDoService.Update(r.Context(), toDoService.UpdateParams{
			ID:          id,
			Name:        req.Name,
			Surname:     req.Surname,
			Patronymic:  req.Patronymic,
			Nationality: req.Nationality,
			Age:         req.Age,
			Gender:      req.Gender,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{ID: id}, http.StatusOK)
	}
}
