package backoffice

import "github.com/rendramakmur/freefolks-fc/helper"

type BackOfficeLoginResponse struct {
	helper.JwtClaims
	AccessToken string `json:"accessToken"`
}
