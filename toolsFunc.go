package mintersdk

import (
	"crypto/ecdsa"
	"encoding/hex"

	//"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/crypto"
	"github.com/MinterTeam/minter-go-node/helpers"
	"github.com/MinterTeam/minter-go-node/rlp"

	// для авторизации/регистрации
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

//////////////////////////////
// Вспомогательные функции
//////////////////////////////

// Генерация новой Seed-фразы
func NewMnemonic() string {
	// Создаёт мнемонику для запоминания или удобный для пользователя seed
	// Мнемоника: это seed фраза
	entropy, _ := bip39.NewEntropy(128) //biteSize должен быть кратен 32 и находиться в пределах включенного диапазона {128, 256}
	Mnemonic, _ := bip39.NewMnemonic(entropy)
	return Mnemonic
}

// Авторизация по Seed-фразе
func AuthMnemonic(seedPhr string) (string, string, error) {
	wallet, err := hdwallet.NewFromMnemonic(seedPhr)
	if err != nil {
		//panic(err)
		return "", "", err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", "", err
	}

	//M+`в нижнем регистре(без видущего нуля)`
	strAdrs := account.Address.String()                                    // 0x512B699Ab21542B8875609593e845818f301903B
	addrss := fmt.Sprintf("M%s", strings.ToLower(strAdrs[1:len(strAdrs)])) // Mx512b699ab21542b8875609593e845818f301903b
	privKeyStr, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", err
	}
	return addrss, privKeyStr, nil
}

// Получение адреса по приватному ключу
func GetAddressPrivateKey(privateKey string) (string, error) {
	privKey2, err := H2ECDSA(privateKey)
	if err != nil {
		return "", err
	}
	addr2 := crypto.PubkeyToAddress(privKey2.PublicKey)
	return addr2.String(), nil
	/* // получаем приватный ключ из объекта ECDSA
	b2 := crypto.FromECDSA(privKey2)
	v2 := encodeHex(b2)*/
}

// конфертирование строки в число с плавающей точкой и коррекция на 18
func cnvStr2Float_18(amntTokenStr string) float32 {
	var fAmntToken float32 = 0.0
	if amntTokenStr != "" {
		fAmntToken64, err := strconv.ParseFloat(amntTokenStr, 64)
		if err != nil {
			panic(err.Error())
		}
		fAmntToken = float32(fAmntToken64 / 1000000000000000000)
	}
	return fAmntToken
}

// HexToECDSA parses a secp256k1 private key.
func H2ECDSA(AccPrivateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(AccPrivateKey)
}

// Encode encodes b as a hex string with 0x prefix.
func encodeHex(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

// Возвращает базовую монету
func GetBaseCoin() types.CoinSymbol {
	return types.GetBaseCoin()
}

// Преобразует строку в монету
func GetStrCoin(coin string) types.CoinSymbol {
	var mntV types.CoinSymbol
	copy(mntV[:], []byte(coin))
	return mntV
}

// Преобразует строку в адрес
func GetStrAddress(addr string) types.Address {
	// Remove Minter wallet prefix and convert hex string to binary
	addrB := types.Hex2Bytes(strings.TrimLeft(addr, "Mx"))

	var adrA types.Address
	copy(adrA[:], addrB)
	return adrA
}

// Целое число в формат pip (18нулей)
func Bip2Pip_i64(value int64) *big.Int {
	return helpers.BipToPip(big.NewInt(value)) // pip в bip(mnt) (!)=косяк, только целочисленные
}

// Число с точкой в формат pip (18нулей)
func Bip2Pip_f64(value float64) *big.Int {
	// FIXME: возможно есть более простая реализация
	mng18 := big.NewInt(1000000000000000) // убрал 000 (3-нуля)
	mng000 := big.NewFloat(1000)          // вот тут 000 (3-нуля)
	amnt := big.NewFloat(value)
	mnFl := big.NewFloat(0).Mul(amnt, mng000)

	amntInt_000, _ := mnFl.Int64()
	var amntBInt big.Int
	amntBInt1 := amntBInt.Mul(big.NewInt(amntInt_000), mng18)

	return amntBInt1
}

// Публичный ключ в массив байтов
func PublicKey2Byte(strPublicKey string) []byte {
	return types.Hex2Bytes(strings.TrimLeft(strPublicKey, "Mp"))
}

func SerializeData(data interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(data)
}
