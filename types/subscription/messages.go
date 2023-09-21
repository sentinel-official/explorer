package subscription

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type MsgCancelRequest struct {
	ID uint64
}

func NewMsgCancelRequest(v bson.M) (*MsgCancelRequest, error) {
	id, err := strconv.ParseUint(v["id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	return &MsgCancelRequest{
		ID: id,
	}, nil
}

type MsgAddQuotaRequest struct {
	From    string
	ID      uint64
	Address string
	Bytes   string
}

func NewMsgAddQuotaRequest(v bson.M) (*MsgAddQuotaRequest, error) {
	id, err := strconv.ParseUint(v["id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	return &MsgAddQuotaRequest{
		From:    v["from"].(string),
		ID:      id,
		Address: v["address"].(string),
		Bytes:   v["bytes"].(string),
	}, nil
}

type MsgUpdateQuotaRequest struct {
	From    string
	ID      uint64
	Address string
	Bytes   string
}

func NewMsgUpdateQuotaRequest(v bson.M) (*MsgUpdateQuotaRequest, error) {
	id, err := strconv.ParseUint(v["id"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	return &MsgUpdateQuotaRequest{
		From:    v["from"].(string),
		ID:      id,
		Address: v["address"].(string),
		Bytes:   v["bytes"].(string),
	}, nil
}
