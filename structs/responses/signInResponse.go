package responses

type UserOriResponseDTO struct {
	StatusCode int
	User       *UsersOri
	StatusDesc string
}

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}

type UserGatewayResponseDTO struct {
	Success    bool
	Result     *UserGateway
	StatusDesc string
}

type UsersGatewayResponseDTO struct {
	Success    bool
	Result     *[]UserGateway
	StatusDesc string
}

type UsersOriResponseDTO struct {
	StatusCode int
	Users      *[]UsersOri
	StatusDesc string
}

type TokenResponseDTO struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int64
}

type LoginDataResponseDTO struct {
	UserType string
	Token    *TokenResponseDTO
}

type LoginTokenResponseDTO struct {
	StatusCode int
	Result     *LoginDataResponseDTO
	StatusDesc string
}

type LoginResponseDTO struct {
	Success    bool
	Result     *LoginDataResponseDTO
	StatusDesc string
}
