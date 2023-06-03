package service

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const pgMethodName = "pg"
const dhMethodName = "dh"

type PgParams struct {
	p           int
	g           int
	nonce       string
	serverNonce string
	messageId   int
}

func (pg PgParams) String() string {
	return fmt.Sprintf("%d %d %s %s %d", pg.p, pg.g, pg.nonce, pg.serverNonce, pg.messageId)
}

type DHParams struct {
	nonce       string
	serverNonce string
	messageId   int
	publicKey   int
}

func getPg(nonce string, requestMessageId string) PgParams {
	// generate the response
	responseMessageId := randomOddInt()
	serverNonce := randomString(20)
	pgResponse := PgParams{23, 5, nonce, serverNonce, responseMessageId}

	// save to cache
	cacheKey := getCacheKey(nonce, serverNonce, pgMethodName)
	cacheData(nil, cacheKey, pgResponse.String(), 20*time.Minute)

	return pgResponse
}

func getDHParams(nonce string, serverNonce string, messageId int, requestPublicKey int) DHParams {
	b := randomInt()
	// get PgParams from cache
	pgCacheKey := getCacheKey(nonce, serverNonce, pgMethodName)
	pgParamsString := getValue(nil, pgCacheKey)
	pgParams, err := getPgParamsFromString(pgParamsString)
	if err != nil {
		return DHParams{}
		// TODO what is the correct thing to do here?
	}

	responsePublicKey := (pgParams.g ^ b) % pgParams.p
	commonKey := (requestPublicKey ^ b) % pgParams.p

	dhCacheKey := getCacheKey(nonce, serverNonce, dhMethodName)
	cacheData(nil, dhCacheKey, string(rune(commonKey)), 20*time.Minute)

	responseMessageId := randomOddInt()
	dhParams := DHParams{
		nonce:       nonce,
		serverNonce: serverNonce,
		messageId:   responseMessageId,
		publicKey:   responsePublicKey,
	}
	return dhParams
}

func getPgParamsFromString(pgParamsString string) (PgParams, error) {
	pgParams := PgParams{}
	_, err := fmt.Sscanf(
		pgParamsString,
		"%d %d %s %s %d",
		&pgParams.p,
		&pgParams.g,
		&pgParams.nonce,
		&pgParams.serverNonce,
		&pgParams.messageId,
	)
	if err != nil {
		fmt.Println(err)
	}
	return pgParams, err

}

func getCacheKey(nonce string, serverNonce string, methodName string) string {
	return methodName + "_" + hashWithSHA1(nonce+serverNonce)
}

func randomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomInt() int {
	return rand.Int()
}

func randomOddInt() int {
	return 2*randomInt() + 1
}

func hashWithSHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	hashInBytes := hash.Sum(nil)
	return string(hashInBytes)
}
