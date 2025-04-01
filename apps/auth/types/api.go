package types

type ENV string

const (
	Production  ENV = "production"
	Development ENV = "development"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
