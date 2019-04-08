package mintersdk

import (
	"crypto/ecdsa"
	"crypto/sha256"

	//"fmt"
	"encoding/binary"
	"math/big"

	c "github.com/MinterTeam/minter-go-node/core/check"
	"github.com/MinterTeam/minter-go-node/crypto"
	"github.com/MinterTeam/minter-go-node/rlp"

	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/crypto/sha3"
)

// Этап 1 - Создание чека
func createCheck(passphrase string, amntMoney float32, coinStr string, privateKey *ecdsa.PrivateKey, nonceID uint64) ([]byte, error) {

	coin := getStrCoin(coinStr)

	passphraseHash := sha256.Sum256([]byte(passphrase))
	passphrasePk, err := crypto.ToECDSA(passphraseHash[:])
	if err != nil {
		//panic(err)
		return []byte{}, err
	}

	nonceIDb := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceIDb, nonceID)

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, nonceID)
	nonceIDb := buf[:n]

	checkValue := bip2pip_f64(float64(amntMoney))

	check := c.Check{
		Nonce:    nonceIDb, //uint64(sdk.GetNonce(AccAddress) + 1), // Уникальный ID чека. Используется для выдачи нескольких одинаковых чеков.
		DueBlock: 999999,   // действителен до блока
		Coin:     coin,
		Value:    checkValue,
	}

	lock, err := crypto.Sign(check.HashWithoutLock().Bytes(), passphrasePk)
	if err != nil {
		//panic(err)
		return []byte{}, err
	}

	check.Lock = big.NewInt(0).SetBytes(lock)

	if err := check.Sign(privateKey); err != nil {
		//panic(err)
		return []byte{}, err
	}

	rawCheck, _ := rlp.EncodeToBytes(check)

	return rawCheck, nil

}

// Этап 2 - Обналичивание чека (точнее proof)
func checkCashingProof(passphrase string, privateKey *ecdsa.PrivateKey) ([65]byte, error) {
	receiverAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	passphraseHash := sha256.Sum256([]byte(passphrase))
	passphrasePk, err := crypto.ToECDSA(passphraseHash[:])
	if err != nil {
		return [65]byte{}, err
	}

	// На адрес получателя receiverAddr
	var senderAddressHash types.Hash
	hw := sha3.NewKeccak256()
	_ = rlp.Encode(hw, []interface{}{
		receiverAddr,
	})
	hw.Sum(senderAddressHash[:0])

	sig, err := crypto.Sign(senderAddressHash.Bytes(), passphrasePk)
	if err != nil {
		return [65]byte{}, err
	}

	proof := [65]byte{}
	copy(proof[:], sig)

	return proof, nil
}
