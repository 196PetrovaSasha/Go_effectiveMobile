package handlers

import (
	"errors"
	"github.com/fir1/rest-api/pkg/erru"

	//"github.com/fir1/rest-api/pkg/erru"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s service) Get() http.HandlerFunc {
	type response struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
		Nationality string `json:"nationality"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
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

		getResponse, err := s.toDoService.Get(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			ID:          getResponse.ID,
			Name:        getResponse.Name,
			Surname:     getResponse.Surname,
			Patronymic:  getResponse.Patronymic,
			Nationality: getResponse.Nationality,
			Age:         getResponse.Age,
			Gender:      getResponse.Gender,
		}, http.StatusOK)
	}
}
