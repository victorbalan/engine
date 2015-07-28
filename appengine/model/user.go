// This file was automatically generated from "../../templates/model/user.got". Do not edit!

package model

const UserKind = "users"

type User struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"-"`
}
