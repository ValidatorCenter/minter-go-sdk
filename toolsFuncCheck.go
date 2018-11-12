package mintersdk

import (
	"crypto/ecdsa"
	"crypto/sha256"

	//"fmt"
	"math/big"

	c "github.com/MinterTeam/minter-go-node/core/check"
	"github.com/MinterTeam/minter-go-node/crypto"
	"github.com/MinterTeam/minter-go-node/rlp"

	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/crypto/sha3"
)

func CreateCheck(passphrase string, amntMoney int64, coinStr string, privateKey *ecdsa.PrivateKey) ([]byte, [65]byte, error) {

	coin := getStrCoin(coinStr)

	passphraseHash := sha256.Sum256([]byte(passphrase))
	passphrasePk, err := crypto.ToECDSA(passphraseHash[:])
	if err != nil {
		//panic(err)
		return []byte{}, [65]byte{}, err
	}

	checkValue := bip2pip_i64(amntMoney)

	check := c.Check{
		Nonce:    0,      //uint64(sdk.GetNonce(AccAddress) + 1), // Уникальный ID чека. Используется для выдачи нескольких одинаковых чеков.
		DueBlock: 999999, // действителен до блока
		Coin:     coin,
		Value:    checkValue,
	}

	lock, err := crypto.Sign(check.HashWithoutLock().Bytes(), passphrasePk)
	if err != nil {
		//panic(err)
		return []byte{}, [65]byte{}, err
	}

	check.Lock = big.NewInt(0).SetBytes(lock)

	if err := check.Sign(privateKey); err != nil {
		//panic(err)
		return []byte{}, [65]byte{}, err
	}

	rawCheck, _ := rlp.EncodeToBytes(check)

	/*receiverPrivateKey, _ := crypto.GenerateKey()
	receiverAddr := crypto.PubkeyToAddress(receiverPrivateKey.PublicKey)*/

	// На адрес получателя receiverAddr
	var senderAddressHash types.Hash
	hw := sha3.NewKeccak256()
	_ = rlp.Encode(hw, []interface{}{
		/*receiverAddr,*/
	})
	hw.Sum(senderAddressHash[:0])

	sig, err := crypto.Sign(senderAddressHash.Bytes(), passphrasePk)
	if err != nil {
		return []byte{}, [65]byte{}, err
	}

	proof := [65]byte{}
	copy(proof[:], sig)

	return rawCheck, proof, nil
}
