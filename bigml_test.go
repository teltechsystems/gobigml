package bigml

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestNewBigML(t *testing.T) {
	Convey("Assuming that env variables are set a bigml object should be created", t, func() {
		bigml, err := NewBigML(os.Getenv("BIGML_USERNAME"), os.Getenv("BIGML_API_KEY"), true)
		So(err, ShouldBeNil)
		So(bigml, ShouldNotBeNil)

		So(bigml.username, ShouldEqual, os.Getenv("BIGML_USERNAME"))
		So(bigml.apiKey, ShouldEqual, os.Getenv("BIGML_API_KEY"))
		So(bigml.client, ShouldEqual, http.DefaultClient)
	})
}

type mockCloser struct {
	io.Reader
}

func (b mockCloser) Close() error { return nil }

type mockTransport struct {
	responses []*http.Response
	offset    int
}

func (t *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	response := t.responses[t.offset%len(t.responses)]
	t.offset += 1
	return response, nil
}

func TestBigMLGetDatasets(t *testing.T) {
	Convey("Assuming a valid mocked response datasets should be returned", t, func() {
		bigml, err := NewBigML(os.Getenv("BIGML_USERNAME"), os.Getenv("BIGML_API_KEY"), true)
		So(err, ShouldBeNil)
		So(bigml, ShouldNotBeNil)

		bigml.client.Transport = &mockTransport{
			responses: []*http.Response{
				&http.Response{Body: mockCloser{strings.NewReader(`{"meta": {"limit": 20, "next": null, "offset": 0, "previous": null, "total_count": 1}, "objects": [{"all_fields": true, "category": 0, "cluster": null, "cluster_status": false, "code": 200, "columns": 16, "created": "2014-08-07T20:09:51.711000", "credits": 0.4763059616088867, "description": "", "dev": true, "field_types": {"categorical": 9, "datetime": 0, "numeric": 7, "preferred": 16, "text": 0, "total": 16}, "locale": "en_US", "missing_numeric_rows": 0, "name": "ultimate_7k's dataset", "number_of_batchcentroids": 0, "number_of_batchpredictions": 0, "number_of_clusters": 0, "number_of_ensembles": 0, "number_of_evaluations": 0, "number_of_models": 1, "number_of_predictions": 502, "objective_field": {"column_number": 15, "datatype": "string", "id": "00000f", "name": "status", "optype": "categorical", "order": 15, "term_analysis": {"enabled": true}}, "price": 0.0, "private": true, "ranges": null, "replacements": null, "resource": "dataset/53e3dd0fffa0440f4300a95c", "rows": 6791, "sample_rates": null, "seeds": null, "shared": false, "size": 499443, "source": "source/53e3dd010af5e83906003c02", "source_status": true, "status": {"bytes": 499443, "code": 5, "elapsed": 1741, "field_errors": [], "message": "The dataset has been created", "row_format_errors": [], "serialized_rows": 6791}, "subscription": false, "tags": [], "term_limit": 32, "updated": "2014-08-07T20:21:21.440000"}]}}`)}},
			},
		}

		datasetResponse, err := bigml.GetDatasets()
		So(err, ShouldBeNil)
		So(datasetResponse, ShouldNotBeNil)
		So(datasetResponse.Meta.Limit, ShouldEqual, 20)
		So(datasetResponse.Meta.Offset, ShouldEqual, 0)
		So(datasetResponse.Meta.TotalCount, ShouldEqual, 1)

		So(len(datasetResponse.Datasets), ShouldEqual, 1)
		So(datasetResponse.Datasets[0].AllFields, ShouldEqual, true)
	})
}
