package mintersdk

const (
	CoinSymbolLength = 10
)

type CoinSymbol [CoinSymbolLength]byte

type SDK struct {
	MnAddress     string // адрес мастер ноды с открытым портом API
	AccAddress    string // адрес кошелька/аккаунта "Mx..."
	AccPrivateKey string // приватный ключ кошелька/аккаунта
}
