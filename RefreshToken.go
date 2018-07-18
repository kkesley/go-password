package password

//RefreshToken holds the refresh token
type RefreshToken struct {
	Secret string `json:"-"`
	Token  string `json:"-"`
	Hash   string `json:"-"`
}
