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

func (ml *BigML) getAuthValues() url.Values {
	return url.Values{
		"username": []string{ml.username},
		"api_key":  []string{ml.apiKey},
	}
}

func (ml *BigML) GetDataset(id string) (*Dataset, error) {
	req, err := http.NewRequest("GET", "/dev/dataset/"+id+"?"+ml.getAuthValues().Encode(), nil)
	if err != nil {
		return nil, err
	}

	res, err := ml.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	ds := &Dataset{}
	if err := json.NewDecoder(res.Body).Decode(ds); err != nil {
		return nil, err
	}

	return ds, nil
}

func (ml *BigML) GetDatasets() (*DatasetResponse, error) {
	req, err := http.NewRequest("GET", "/dev/dataset?"+ml.getAuthValues().Encode(), nil)
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
