package service

import (
	"fmt"
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type PgParams struct {
	p           int
	g           int
	nonce       string
	serverNonce string
	messageId   int
}

type DHParams struct {
	nonce       string
	serverNonce string
	messageId   int
	publicKey   int
}

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomInt() int {
	return rand.Int()
}

func getPg(nonce string, requestMessageId string) PgParams {
	responseMessageId := 2*randomInt() + 1
	serverNonce := randomString(20)
	pgResponse := PgParams{23, 5, nonce, serverNonce, responseMessageId}
	// TODO cache response
	return pgResponse
}

func getDHParams(nonce string, serverNonce string, messageId int, requestPublicKey int) DHParams {
	b := randomInt()
	// TODO get PgParams from cache
	pgParams := PgParams{23, 5, nonce, serverNonce, messageId}

	responsePublicKey := (pgParams.g ^ b) % pgParams.p

	commonKey := (requestPublicKey ^ b) % pgParams.p
	fmt.Println(commonKey)
	// TODO cache common key

	responseMessageId := 2*randomInt() + 1
	dhParams := DHParams{
		nonce:       nonce,
		serverNonce: serverNonce,
		messageId:   responseMessageId,
		publicKey:   responsePublicKey,
	}
	return dhParams
}
