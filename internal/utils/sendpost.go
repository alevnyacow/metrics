package utils

import (
	"net/http"
)

func SendPost(url string) (requestError error, response *http.Response, responseError error) {
	request, requestErr := http.NewRequest("POST", url, nil)
	if requestErr != nil {
		requestError = requestErr

		return
	}

	request.Header.Set("Content-Type", "text/plain")
	client := http.Client{}
	response, responseError = client.Do(request)

	return
}
