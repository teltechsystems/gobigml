package bigml

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type BigML struct {
	username string
	apiKey   string
	client   *http.Client
}

func (ml *BigML) GetDatasets() (*DatasetResponse, error) {
	values := url.Values{
		"username": []string{ml.username},
		"api_key":  []string{ml.apiKey},
	}

	req, err := http.NewRequest("GET", "/dev/dataset?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}

	res, err := ml.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dr := &DatasetResponse{}
	if err := json.NewDecoder(res.Body).Decode(dr); err != nil {
		return nil, err
	}

	return dr, nil
}

func NewBigML(username, apiKey string, devMode bool) (*BigML, error) {
	return &BigML{
		username: username,
		apiKey:   apiKey,
		client:   http.DefaultClient,
	}, nil
}
