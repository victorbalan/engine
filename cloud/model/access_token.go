// This file was automatically generated from "../../templates/model/access_token.got". Do not edit!

package model

import "time"

const AccessTokenKind = "accesstokens"

type AccessToken struct {
	Value        string    `json:"value"`
	Scopes       []string  `json:"scopes"`
	Description  string    `json:"description"`
	Creation     time.Time `json:"creation"`
	Modification time.Time `json:"modification"`
	Expiry       time.Time `json:"expiry"`
}
