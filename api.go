package pokego

import (
	"fmt"
	"io"
	"net/http"
)

func callApi(endpoint string) ([]byte, error) {
	var url = fmt.Sprintf("https://pokeapi.co/api/v2/%v", endpoint)

	req, err := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)

	return body, err
}
