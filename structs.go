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

const (
	TX_SendData             int = iota + 1 //1
	TX_SellCoinData                        //2
	TX_SellAllCoinData                     //3
	TX_BuyCoinData                         //4
	TX_CreateCoinData                      //5
	TX_DeclareCandidacyData                //6
	TX_DelegateDate                        //7
	TX_UnbondData                          //8
	TX_RedeemCheckData                     //9
	TX_SetCandidateOnData                  //10
	TX_SetCandidateOffData                 //11
	TX_CreateMultisigData                  //12
)
