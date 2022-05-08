package codeforcesv2

import (
	"crypto/sha512"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

const (
	//sum of AC number
	cfPassNumber = "codeforces_Person_sum_problem"
)

//got public problems passed api
func statusWithoutKey(userName string) string {
	wget := "https://codeforces.com/api/user.status?handle=" + userName
	return wget
}

//got including private problems api
func statusWithKey(userName string, key string, secret string) string {

	//watch https://codeforces.com/apiHelp
	ti := int(time.Now().Unix())
	sig := "123456"
	secret = "#" + secret
	str := "/user.status?apiKey=" + key + "&handle=" + userName + "&time=" + strconv.Itoa(ti)
	hash := []byte(sig + str + secret)
	sha := sha512.New()
	io.WriteString(sha, string(hash))
	bw := sha.Sum(nil)
	hashSign := hex.EncodeToString(bw)
	wget := "https://codeforces.com/api" + str + "&apiSig=" + sig + hashSign

	log.Println(wget)
	return wget
}
