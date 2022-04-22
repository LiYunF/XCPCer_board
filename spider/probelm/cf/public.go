package cf

import "fmt"

func getPersonSubmissionsInit(uid string) string {
	return "https://codeforces.com/api/user.status?handle=" + uid
}
func getPersonSubmissions(uid string, from int, count int) string {
	//uid 	(Required)	Codeforces user handle.
	//	from	1-based index of the first submission to return.
	//count		Number of returned submissions.
	return fmt.Sprintf("https://codeforces.com/api/user.status?handle=%v&from=%v&count=%v", uid, from, count)
}
