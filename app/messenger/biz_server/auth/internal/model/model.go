package model

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"io"
	"strconv"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
)

type PhoneCodeTransaction struct {
	AuthKeyId             int64  `json:"auth_key_id"`
	SessionId             int64  `json:"session_id"`
	PhoneNumber           string `json:"phone_number"`
	PhoneNumberRegistered bool   `json:"phone_number_registered"`
	PhoneCode             string `json:"phone_code"`
	PhoneCodeHash         string `json:"phone_code_hash"`
	PhoneCodeExpired      int32  `json:"phone_code_expired"`
	PhoneCodeExtraData    string `json:"phone_code_extra_data"`
	SentCodeType          int    `json:"sent_code_type"`
	FlashCallPattern      string `json:"flash_call_pattern"`
	NextCodeType          int    `json:"next_code_type"`
	State                 int    `json:"state"`
}

func (m *PhoneCodeTransaction) ToAuthSentCode() *mtproto.Auth_SentCode {
	authSentCode := mtproto.MakeTLAuthSentCode(&mtproto.Auth_SentCode{
		PhoneRegistered: m.PhoneNumberRegistered,
		Type:            makeAuthSentCodeType(m.SentCodeType, len(m.PhoneCode), m.FlashCallPattern),
		PhoneCodeHash:   m.PhoneCodeHash,
		NextType:        makeAuthCodeType(m.NextCodeType),
		Timeout:         &types.Int32Value{Value: 60},
	}).To_Auth_SentCode()
	if m.SentCodeType == CodeTypeApp {
		authSentCode.Timeout = nil
	}
	return authSentCode
}

const (
	QRCodeStateNew      = 1
	QRCodeStateAccepted = 2
	QRCodeStateSuccess  = 3
)

type QRCodeTransaction struct {
	AuthKeyId int64  `json:"auth_key_id"`
	ServerId  string `json:"server_id"`
	SessionId int64  `json:"session_id"`
	ApiId     int32  `json:"api_id"`
	ApiHash   string `json:"api_hash"`
	CodeHash  string `json:"code_hash"`
	ExpireAt  int64  `json:"expire_at"`
	UserId    int32  `json:"user_id"`
	State     int    `json:"state"`
}

func (m *QRCodeTransaction) Token() []byte {
	token := make([]byte, 8, 24)
	binary.BigEndian.PutUint64(token, uint64(m.AuthKeyId))
	m2 := md5.New()
	io.WriteString(m2, strconv.Itoa(int(m.AuthKeyId)))
	io.WriteString(m2, m.CodeHash)
	io.WriteString(m2, strconv.Itoa(int(m.ExpireAt)))
	return m2.Sum(token)
}

func (m *QRCodeTransaction) CheckByToken(token []byte) bool {
	return bytes.Equal(m.Token(), token)
}
