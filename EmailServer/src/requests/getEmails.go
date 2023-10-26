package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dhliv/EmailServer/src/constants"
	"github.com/dhliv/EmailServer/src/models"
)

const PAGE_SIZE int = 10

func GetEmails(page int) (*models.Emails, error) {
	url := fmt.Sprintf("%v/api/%v/_search", constants.API_URL, constants.API_INDEX)
	query := fmt.Sprintf(
		`{
			"query": {
					"match_all": {}
			},
			"max_results": %v,
			"from": %v,
			"sort_fields": ["_id"]
	}`, PAGE_SIZE, page)

	req, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(constants.USER_NAME, constants.USER_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	emails := models.EmailsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&emails)
	if err != nil {
		return nil, err
	}

	return &emails.Hits, nil
}