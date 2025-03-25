package utils

import (
	"fmt"
	"net/http"
)

func SendPost(url string) (requestError error, response *http.Response, responseError error) {
	fmt.Println("Sent post to " + url)
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
