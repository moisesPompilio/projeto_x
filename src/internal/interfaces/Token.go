package interfaces

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
