package netCommon

type IHttpClient interface {
	Update() error
	// TODO: add return
	AsyncRequest(
		method string,
		ip string,
		port uint16,
		url string,
		params map[string]string,
		cookies []string,
		body string,
	)
}
