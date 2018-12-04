# Minter Golang SDK
![](minter-go.svg)

## About
Minter Blockchain Golang SDK https://minter.network

Actual for Minter version 0.5.0.

* [Installation](#installing)
* [Updating](#updating)
* [Minter SDK](#using-minter)
	- [GetBalance](#getbalance)
	- [GetNonce](#getnonce)
	- [GetStatus](#getstatus)
	- [GetValidators](#getvalidators)
	- [GetValidatorsBlock](#getvalidatorsblock)
	- [EstimateCoinBuy](#estimatecoinbuy)
	- [EstimateCoinSell](#estimatecoinsell)
	- [GetCoinInfo](#getcoininfo)
	- [GetBlock](#getblock)
	- [GetTransaction](#gettransaction)
	- [GetCandidate](#getcandidate)
	- [GetCandidates](#getcandidates)
	- [NewMnemonic](#newmnemonic)
	- [AuthMnemonic](#authmnemonic)
	- [GetAddressPrivateKey](#getaddressprivatekey)
	- [GetVAddressPubKey](#getvaddresspubkey)
	- [SendCoin](#example-2)
	- [MultiSendCoin](#example-3)
	- [SellCoinTx](#example-4)
	- [SellAllCoin](#example-5)
	- [BuyCoinTx](#example-6)
	- [CreateCoin](#example-7)
	- [DeclareCandidacy](#example-8)
	- [Delegate](#example-9)
	- [SetCandidate](#example-10)
  

## Installing

```bash
go get github.com/ValidatorCenter/minter-go-sdk
```

## Updating

```bash
go get -u github.com/ValidatorCenter/minter-go-sdk
```

## Using Minter

Create MinterSDK instance.

```golang
import m "github.com/ValidatorCenter/minter-go-sdk"

sdk := m.SDK{
		MnAddress: "http://156.123.34.5:8841", // example of a node url
}
```

Structures for receiving data from the blockchain already have tags: json, bson and gorm. This means that you can immediately write them correctly to a database that uses one of these tags.

### GetBalance

Returns coins list and balance of an address.

``
GetBalance("Mx...MinterAddress" string): map[string]string, error
``

###### Example

```golang
blnc, err := sdk.GetBalance("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")

// result: {MTN: 1000000, TESTCOIN: 2000000}
```

### GetNonce

Returns current nonce of an address.

``
GetNonce("Mx...MinterAddress" string): int, error
``

###### Example

```golang
lastNmb, err := sdk.GetNonce("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")

// 5
```

### GetStatus

Returns node status info.

``
GetStatus(): struct, error
``

### GetValidators

Returns list of active validators.

``
GetValidators(): struct, error
``

### GetValidatorsBlock

Returns a list of validators of a block by its number.

``
GetValidatorsBlock("blockNumber" int): struct, error
``

### EstimateCoinBuy

Return estimate of buy coin transaction.

``
EstimateCoinBuy("coinToBuy" string, "coinToSell" coinToBuy, "valueToBuy" int64): struct, error
``

### EstimateCoinSell

Return estimate of sell coin transaction.

``
EstimateCoinSell("coinToSell" string, "coinToBuy" string, "valueToSell" int64): struct, error
``

### GetCoinInfo

Returns information about coin.

``
GetCoinInfo("COIN_SYMBOL" string): struct, error
``

### GetBlock

Returns block data at given height.

``
GetBlock("height" int): struct, error
``

### GetTransaction

Returns transaction info.

``
GetTransaction("Mt...hash" string): struct, error
``

### GetCandidate

Returns candidate’s info by provided public_key. It will respond with 404 code if candidate is not found.

``
GetCandidate("Mp..." string): struct, error
``

### GetCandidates

Returns list of candidates.

``
GetCandidates(): struct, error
``

### NewMnemonic

Returns new seed-phrase.

``
NewMnemonic(): string, error
``

### AuthMnemonic

Authorization by seed-phrase.

``
AuthMnemonic("seed-phrase" string): "address" string, "private-key" string, error
``

### GetAddressPrivateKey

Returns address of the wallet by provided private key.

``
GetAddressPrivateKey("private-key" string): "Mx..." string, error
``

### GetVAddressPubKey

Returns validator-address by provided public key.

``
GetVAddressPubKey("Mp..." string): string
``

### Sign transaction

Returns a signed tx

###### Example

* Sign the <b>SendCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	sndDt := m.TxSendCoinData{
		Coin:      "MNT",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7",
		Value:     10,
		GasCoin:   "MNT",
		GasPrice:  1,
	}

	resHash, err := sdk.TxSendCoin(&sndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example

* Sign the <b>MultiSendCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "...",
	}

	cntList := []m.TxOneSendCoinData{}

	// First address
	cntList = append(cntList, TxOneSendCoinData{
		Coin:      "MNT",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7", //Кому переводим
		Value:     10,
	})

	// Second address
	cntList = append(cntList, TxOneSendCoinData{
		Coin:      "VALIDATOR",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7", //Кому переводим
		Value:     16,
	})

	mSndDt := m.TxMultiSendCoinData{
		List:     cntList,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxMultiSendCoin(&mSndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>SellCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	slDt := m.TxSellCoinData{
		CoinToSell:  "MNT",
		CoinToBuy:   "ABCDEF24",
		ValueToSell: 10,
		GasCoin:     "MNT",
		GasPrice:    1,
	}

	resHash, err := sdk.TxSellCoin(&slDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>SellAllCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	slAllDt := m.TxSellAllCoinData{
		CoinToSell: "MNT",
		CoinToBuy:  "ABCDEF24",
		GasCoin:    "MNT",
		GasPrice:   1,
	}

	resHash, err := sdk.TxSellAllCoin(&slAllDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>BuyCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	buyDt := m.TxBuyCoinData{
		CoinToSell: "ABCDEF23",
		CoinToBuy:  "MNT",
		ValueToBuy: 1,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxBuyCoin(&buyDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}

```

###### Example
* Sign the <b>CreateCoin</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	creatDt := m.TxCreateCoinData{
		Name:                 "Test coin 24",
		Symbol:               "ABCDEF24",
		InitialAmount:        100,
		InitialReserve:       100,
		ConstantReserveRatio: 50,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxCreateCoin(&creatDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>DeclareCandidacy</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	declDt := m.TxDeclareCandidacyData{
		PubKey:     "Mp2891198c692c351bc55ac60a03c82649fa920f7ad20bd290a0c4e774e916e9de", // "Mp....",
		Commission: 10,                                                                   // 10%
		Coin:       "MNT",
		Stake:      100,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxDeclareCandidacy(&declDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)
}
```

###### Example
* Sign the <b>Delegate</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	delegDt := m.TxDelegateData{
		Coin:     "MNT",
		PubKey:   "Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610",
		Stake:    100,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxDelegate(&delegDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>SetCandidate</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	sndDt := m.TxSetCandidateData{
		PubKey:   "Mp2891198c692c351bc55ac60a03c82649fa920f7ad20bd290a0c4e774e916e9de",
		Activate: true, //true-"on", false-"off"
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxSetCandidate(&sndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```

###### Example
* Sign the <b>Unbound</b> transaction.

```golang
package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "your private key",
	}

	unbDt := m.TxUnbondData{
		PubKey:   "Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610",
		Coin:     "MNT",
		Value:    10,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxUnbond(&unbDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
```
