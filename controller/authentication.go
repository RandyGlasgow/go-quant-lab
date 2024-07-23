package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/authorizerdev/authorizer-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vec-search/lib/auth/auth_service"
	"github.com/vec-search/lib/http/http_utils"
)

func AuthenticationRoutes(r chi.Router) {

	r.Route("/auth", func(r chi.Router) {

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {

			requestJson := authorizer.LoginInput{}
			if err := http_utils.ExtractJsonFromRequestBody(r, &requestJson); err != nil {
				http_utils.HttpCustomError(w, errors.New("error reading request body"))
				return
			}

			fmt.Printf("requestJson: %+v\n", requestJson)

			res, err := auth_service.AuthClient.Login(&requestJson)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}
			// set the vec-search auth token
			header := w.Header()
			header.Set("Authorization", *res.AccessToken)
			header.Set("ID_Token", *res.IdToken)

			render.JSON(w, r, res)
		})

		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {

			requestJson := authorizer.SignUpInput{}

			if err := http_utils.ExtractJsonFromRequestBody(r, &requestJson); err != nil {
				http_utils.HttpCustomError(w, errors.New("error reading request body"))
				return
			}

			res, err := auth_service.AuthClient.SignUp(&requestJson)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			render.JSON(w, r, res)
		})
	})
}
