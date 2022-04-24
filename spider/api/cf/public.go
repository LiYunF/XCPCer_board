package cf

import (
	"crypto/sha512"
	"fmt"
	"time"
)

func getapikey(apikey string, secret string) {
	apikey = "&apiKey=" + apikey
	time := time.Now() //获取当前时间

	fmt.Println(time)
	//time = "&time=" + strconv.Itoa(time)
}
func getPersonSubmissionsInit(uid string) string {

	return "https://codeforces.com/api/user.status?handle=" + uid + fmt.Sprintf("%x", sha512.Sum512([]byte(uid)))
}
func getPersonSubmissions(uid string, from int, count int) string {
	//uid 	(Required)	Codeforces user handle.
	//	from	1-based index of the first submission to return.
	//count		Number of returned submissions.
	return fmt.Sprintf("https://codeforces.com/api/user.status?handle=%v&from=%v&count=%v", uid, from, count)
}
