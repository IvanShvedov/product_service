package server

import (
	"encoding/json"
	"net/http"
	"product-service/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) create_product() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product := model.CreateProductDTO{}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			s.error(w, r, 500, err)
			return
		}
		service := s.service.Product()
		res, err := service.Create(product)
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		s.respond(w, r, 200, res)
	}
}

func (s *Server) delete_product() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := request{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.error(w, r, 500, err)
			return
		}
		service := s.service.Product()
		res, err := service.Delete(request.Id)
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		s.respond(w, r, 200, res)
	}
}

func (s *Server) update_product() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product := model.UpdateProductDTO{}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			s.error(w, r, 500, err)
			return
		}
		service := s.service.Product()
		res, err := service.Update(product)
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		s.respond(w, r, 200, res)
	}
}

func (s *Server) fetch_one_product() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		service := s.service.Product()
		res, err := service.FetchOne(id)
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		s.respond(w, r, 200, res)
	}
}

func (s *Server) fetch_all_product() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := s.service.Product()
		res, err := service.FetchAll()
		if err != nil {
			s.error(w, r, 500, err)
			return
		}
		s.respond(w, r, 200, res)
	}
}
