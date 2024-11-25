package models

type UsersResponse struct {
	Status  int        `json:"status"`
	Code    string     `json:"code"`
	Message string     `json:"message,omitempty"`
	Data    *UsersData `json:"data,omitempty"`
}
type UsersRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UsersData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type PokemonResponse struct {
	Status  int                     `json:"status"`
	Code    string                  `json:"code"`
	Message string                  `json:"message,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty"`
}
