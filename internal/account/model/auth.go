/*
    Package model  for 'Auth' DTO's 
*/
package model

import "time"

type AuthLoginDTO struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthLoginResponse struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    TransmissionKey string  `json:"transmission_key"`
}

type TokenDetailsDTO struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    AtExpiresTime   time.Time
    RtExpiresTime   time.Time
    TransmissionKey string  `json:"transmission_key"`
}

