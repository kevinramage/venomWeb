package api

//https://chromium.googlesource.com/chromium/src/+/master/docs/chromedriver_status.md

type NavigateRequest struct {
	Url string `json:"url"`
}
type NavigateResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}
