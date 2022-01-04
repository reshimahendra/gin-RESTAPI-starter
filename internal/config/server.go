package config

type ServerConfiguration struct {
    Port                       string
    SecretKey                  string
    AccessTokenExpireDuration  int64
    RefreshTokenExpireDuration int64
    LimitCountPerRequest       float64
}
