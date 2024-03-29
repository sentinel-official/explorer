package models

import (
	"encoding/json"
	"maps"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/bytes"
	tmtypes "github.com/tendermint/tendermint/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/sentinel-official/explorer/types"
	"github.com/sentinel-official/explorer/utils"
)

type Message struct {
	Data bson.M `json:"data,omitempty" bson:"data"`
	Type string `json:"type,omitempty" bson:"type"`
}

func NewMessage(v sdk.Msg) *Message {
	item := &Message{
		Data: bson.M{},
		Type: utils.MsgTypeURL(v),
	}

	buf, err := types.EncCfg.Codec.MarshalJSON(v)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(buf, &item.Data); err != nil {
		panic(err)
	}

	return item
}

type Messages []*Message

func NewMessages(v []sdk.Msg) Messages {
	items := make(Messages, 0, len(v))
	for _, item := range v {
		items = append(items, NewMessage(item))
	}

	return items
}

func (m Messages) WithAuthzMsgExecMessages() (items Messages) {
	for i := 0; i < len(m); i++ {
		items = append(items, m[i])
		if strings.Contains(m[i].Type, "cosmos.authz") && strings.Contains(m[i].Type, "MsgExec") {
			msgs := m[i].Data["msgs"].(bson.A)
			for j := 0; j < len(msgs); j++ {
				item := &Message{
					Data: func() bson.M {
						m := make(bson.M)
						maps.Copy(m, msgs[j].(bson.M))
						delete(m, "@type")

						return m
					}(),
					Type: msgs[j].(bson.M)["@type"].(string),
				}

				items = append(items, item)
			}
		}
	}

	return items
}

type TxResult struct {
	Codespace string       `json:"codespace,omitempty" bson:"codespace"`
	Code      uint32       `json:"code,omitempty" bson:"code"`
	Events    types.Events `json:"events,omitempty" bson:"events"`
	GasUsed   int64        `json:"gas_used,omitempty" bson:"gas_used"`
	GasWanted int64        `json:"gas_wanted,omitempty" bson:"gas_wanted"`
	Info      string       `json:"info,omitempty" bson:"info"`
	Log       string       `json:"logs,omitempty" bson:"log"`
}

func NewTxResult(v *abcitypes.ResponseDeliverTx) *TxResult {
	return &TxResult{
		Codespace: v.Codespace,
		Code:      v.Code,
		Events:    types.NewEventsFromABCIEvents(v.Events),
		GasUsed:   v.GasUsed,
		GasWanted: v.GasWanted,
		Info:      v.Info,
		Log:       types.Replacer.Replace(v.Log),
	}
}

type TxSignerInfo struct {
	Address   string      `json:"address,omitempty" bson:"address"`
	Mode      interface{} `json:"mode,omitempty" bson:"mode"`
	PublicKey string      `json:"public_key,omitempty" bson:"public_key"`
	Sequence  uint64      `json:"sequence,omitempty" bson:"sequence"`
	Signature string      `json:"signature,omitempty" bson:"signature"`
}

type TxSignerInfos []*TxSignerInfo

func NewTxSignerInfosFromTx(v authsigning.Tx) TxSignerInfos {
	signatures, err := v.GetSignaturesV2()
	if err != nil {
		panic(err)
	}

	var (
		signers = v.GetSigners()
		items   = make(TxSignerInfos, 0, len(signers))
	)

	for i := 0; i < len(signers); i++ {
		modeInfo, signature := authtx.SignatureDataToModeInfoAndSig(signatures[i].Data)

		items = append(
			items,
			&TxSignerInfo{
				Address:   signers[i].String(),
				Mode:      modeInfo,
				PublicKey: bytes.HexBytes(signatures[i].PubKey.Bytes()).String(),
				Sequence:  signatures[i].Sequence,
				Signature: bytes.HexBytes(signature).String(),
			},
		)
	}

	return items
}

type Tx struct {
	Fee           types.Coins   `json:"fee,omitempty" bson:"fee"`
	GasLimit      uint64        `json:"gas_limit,omitempty" bson:"gas_limit"`
	Granter       string        `json:"granter,omitempty" bson:"granter"`
	Hash          string        `json:"hash,omitempty" bson:"hash"`
	Height        int64         `json:"height,omitempty" bson:"height"`
	Index         int           `json:"index,omitempty" bson:"index"`
	Memo          string        `json:"memo,omitempty" bson:"memo"`
	Messages      Messages      `json:"messages,omitempty" bson:"messages"`
	Payer         string        `json:"payer,omitempty" bson:"payer"`
	Result        *TxResult     `json:"result,omitempty" bson:"result"`
	SignerInfos   TxSignerInfos `json:"signer_infos,omitempty" bson:"signer_infos"`
	TimeoutHeight uint64        `json:"timeout_height,omitempty" bson:"timeout_height"`
	Timestamp     time.Time     `json:"timestamp,omitempty" bson:"timestamp"`
}

func NewTx(v tmtypes.Tx) *Tx {
	t, err := types.EncCfg.TxConfig.TxDecoder()(v)
	if err != nil {
		return &Tx{
			Hash: bytes.HexBytes(v.Hash()).String(),
		}
	}

	tx := t.(authsigning.Tx)
	return &Tx{
		Fee:           types.NewCoins(tx.GetFee()),
		GasLimit:      tx.GetGas(),
		Granter:       tx.FeeGranter().String(),
		Hash:          bytes.HexBytes(v.Hash()).String(),
		Height:        0,
		Index:         0,
		Memo:          tx.GetMemo(),
		Messages:      NewMessages(tx.GetMsgs()),
		Payer:         tx.FeePayer().String(),
		SignerInfos:   NewTxSignerInfosFromTx(tx),
		TimeoutHeight: tx.GetTimeoutHeight(),
		Timestamp:     time.Time{},
	}
}

func (t *Tx) String() string {
	return utils.MustMarshalIndentToString(t)
}

func (t *Tx) WithHeight(v int64) *Tx                        { t.Height = v; return t }
func (t *Tx) WithIndex(v int) *Tx                           { t.Index = v; return t }
func (t *Tx) WithResult(v *abcitypes.ResponseDeliverTx) *Tx { t.Result = NewTxResult(v); return t }
func (t *Tx) WithTimestamp(v time.Time) *Tx                 { t.Timestamp = v; return t }
