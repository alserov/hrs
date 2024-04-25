package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

var (
	sdAddr = os.Getenv("SERVICE_DISCOVERY_ADDR")
)

func Discover(group string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, sdAddr, nil)
	if err != nil {
		return nil, NewError(ErrInternal, err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, NewError(ErrInternal, err.Error())
	}
	defer res.Body.Close()

	var addrs []string
	if err = json.NewDecoder(res.Body).Decode(&addrs); err != nil {
		return nil, NewError(ErrInternal, err.Error())
	}

	return addrs, nil
}
