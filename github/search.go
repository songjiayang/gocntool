package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SearchResult struct {
	TotalCount int `json:"total_count"`
}

func Search(q string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/search/code?q=%s+in:file+repo:gocn/news", q)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result SearchResult

	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}

	return result.TotalCount > 0, nil
}
