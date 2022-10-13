package service

import autoredial "github.com/tk42/auto-redial/gen/sqlc"

type ServiceServer struct {
	db *autoredial.Queries
}
