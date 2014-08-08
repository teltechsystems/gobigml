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

func TestBigMLGetAuthValues(t *testing.T) {
	Convey("A simple auth string should be generated", t, func() {
		bigml, err := NewBigML(os.Getenv("BIGML_USERNAME"), os.Getenv("BIGML_API_KEY"), true)
		So(err, ShouldBeNil)
		So(bigml, ShouldNotBeNil)

		So(bigml.getAuthValues().Encode(), ShouldEqual, "api_key="+os.Getenv("BIGML_API_KEY")+"&username="+os.Getenv("BIGML_USERNAME"))
	})
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

func TestBigMLGetDataset(t *testing.T) {
	Convey("Assuming a valid mocked response a single dataset should be returned", t, func() {
		bigml, err := NewBigML(os.Getenv("BIGML_USERNAME"), os.Getenv("BIGML_API_KEY"), true)
		So(err, ShouldBeNil)
		So(bigml, ShouldNotBeNil)

		bigml.client.Transport = &mockTransport{
			responses: []*http.Response{
				&http.Response{Body: mockCloser{strings.NewReader(`{"all_fields": true, "category": 0, "cluster": null, "cluster_status": false, "code": 200, "columns": 16, "created": "2014-08-07T20:09:51.711000", "credits": 0.4763059616088867, "description": "", "dev": true, "download": {"code": 0, "excluded_input_fields": [], "header": true, "input_fields": [], "message": "", "preview": [], "separator": ","}, "excluded_fields": [], "field_types": {"categorical": 9, "datetime": 0, "numeric": 7, "preferred": 16, "text": 0, "total": 16}, "fields": {"000000": {"column_number": 0, "datatype": "string", "name": "cnam_via_sms", "optype": "categorical", "order": 0, "preferred": true, "summary": {"categories": [["1", 4718], ["0", 2073]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "000001": {"column_number": 1, "datatype": "int16", "name": "total_calls", "optype": "numeric", "order": 1, "preferred": true, "summary": {"bins": [[24.35281, 3543], [110.40746, 1367], [190.13164, 904], [279.51697, 383], [349.13433, 134], [401.36957, 92], [460.07273, 110], [533.22619, 84], [607.60714, 28], [661.96, 25], [727.79167, 24], [807.70588, 17], [890.73684, 19], [967, 13], [1018.5, 8], [1093.33333, 3], [1147, 1], [1215.33333, 3], [1294.75, 4], [1407.125, 8], [1487, 1], [1589.33333, 3], [1672, 2], [1809.5, 2], [2115.5, 2], [2195, 2], [2424, 2], [2694.5, 2], [3843, 1], [3961, 1], [4064, 2], [8076, 1]], "maximum": 8076, "mean": 129.32528, "median": 71.83591, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [1.1373, 12.85714, 33.30437, 56.0058, 88.1038, 125.41831, 178.45626, 283.9141], "standard_deviation": 229.58477, "sum": 878248, "sum_squares": 471474924, "variance": 52709.16825}}, "000002": {"column_number": 2, "datatype": "int16", "name": "total_unmasked", "optype": "numeric", "order": 2, "preferred": true, "summary": {"bins": [[3.04456, 5880], [17.67325, 456], [28.66292, 178], [38.6988, 83], [46.60417, 48], [59.18605, 43], [69.13333, 15], [78.36364, 11], [85.81818, 11], [95.375, 8], [103.54545, 11], [112.33333, 6], [122.25, 4], [143.5, 2], [155.6, 5], [166, 4], [176.33333, 3], [187, 2], [201.5, 2], [217.25, 4], [234, 1], [244, 1], [252, 1], [267, 2], [322, 3], [356, 1], [435, 1], [525, 1], [640, 1], [674, 1], [888, 1], [2610, 1]], "maximum": 2610, "mean": 8.42689, "median": 2.83814, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [0.58788, 2.83814, 7.71543], "standard_deviation": 40.72165, "sum": 57227, "sum_squares": 11741781, "variance": 1658.25264}}, "000003": {"column_number": 3, "datatype": "int16", "name": "total_blacklisted", "optype": "numeric", "order": 3, "preferred": true, "summary": {"bins": [[0.21415, 6206], [9.29795, 292], [20.37374, 99], [29.06667, 30], [34.51852, 27], [42.45455, 33], [49.4, 15], [57.52381, 21], [66.375, 16], [73, 5], [79.5, 6], [88, 8], [95.5, 2], [103.5, 4], [111, 2], [117.5, 2], [126, 2], [133.75, 4], [140, 2], [166.5, 2], [182, 1], [202, 1], [214, 1], [240, 1], [260, 1], [277, 1], [331, 1], [454, 1], [459.5, 2], [531, 1], [660, 1], [3669, 1]], "maximum": 3669, "mean": 3.56707, "median": 0.23839, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [0.23839], "standard_deviation": 48.47919, "sum": 24224, "sum_squares": 16044480, "variance": 2350.2314}}, "000004": {"column_number": 4, "datatype": "int16", "name": "first_call", "optype": "numeric", "order": 4, "preferred": true, "summary": {"bins": [[1.88903, 4884], [25.15174, 402], [47.85714, 119], [68.68182, 66], [88.62791, 43], [101.5, 4], [114.5, 12], [131.05556, 18], [157.72222, 18], [186.4, 10], [207.33333, 3], [220.33333, 6], [237.2, 5], [258, 2], [279, 5], [301.71429, 7], [334.33333, 3], [366, 4], [380, 2], [400, 1], [440, 1], [457, 1], [479, 1], [498, 3], [544, 1], [561, 1], [575.66667, 3], [609, 1], [626, 2], [650, 3], [688, 2], [8760, 1158]], "maximum": 8760, "mean": 1503.1611, "median": 3.20714, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [0.74602, 10.25582], "standard_deviation": 3290.77009, "sum": 10207967, "sum_squares": 88874268133, "variance": 10829167.787}}, "000005": {"column_number": 5, "datatype": "int16", "name": "last_call", "optype": "numeric", "order": 5, "preferred": true, "summary": {"bins": [[8.40078, 257], [43.33684, 95], [73.80328, 61], [92.68085, 47], [113.84746, 59], [143.42105, 57], [170.775, 40], [195.40426, 47], [217.77778, 54], [243.27083, 48], [266.37209, 43], [290.0625, 48], [313.0303, 33], [334.54167, 24], [356.84906, 53], [381.2, 30], [406.12676, 71], [430.225, 40], [454.86275, 51], [477.08333, 48], [506.01724, 58], [530.58974, 39], [550.6875, 64], [573.32653, 49], [596.73418, 79], [619.7191, 89], [644.16556, 151], [667.86093, 302], [692.97886, 615], [715.94852, 1282], [738.47675, 1699], [8760, 1158]], "maximum": 8760, "mean": 1986.91268, "median": 717.67032, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [183.87785, 505.73351, 666.79726, 697.68742, 714.15191, 721.75759, 734.21459, 740.49403, 4382.09626, 8494.92713], "standard_deviation": 3077.99087, "sum": 13493124, "sum_squares": 91138307916, "variance": 9474027.80077}}, "000006": {"column_number": 6, "datatype": "double", "name": "total_spent", "optype": "numeric", "order": 6, "preferred": true, "summary": {"bins": [[0.00112, 1331], [4.98714, 35], [10.34062, 32], [18.71, 7], [24.94941, 3593], [29.9, 75], [34.93297, 64], [39.85, 3], [44.85, 5], [49.9004, 1458], [54.89867, 15], [59.71278, 18], [69.8, 5], [74.81528, 72], [79.9275, 8], [84.8, 2], [94.93333, 3], [99.80462, 26], [104.75, 1], [109.75, 1], [120.35, 1], [124.75, 3], [129.0875, 8], [145.3, 2], [153.98375, 8], [178.9, 1], [215.75, 2], [239.2625, 4], [247.38, 1], [264.35, 2], [272.33, 1], [289.3, 4]], "maximum": 289.3, "mean": 27.42129, "median": 25.22419, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [0.01588, 5.42034, 23.92421, 24.58263, 25.22419, 26.34503, 28.02124, 49.85705, 49.92521], "standard_deviation": 21.80775, "sum": 186218.01, "sum_squares": 8335514.2043, "variance": 475.57811}}, "000007": {"column_number": 7, "datatype": "string", "name": "voicemail_carrier", "optype": "categorical", "order": 7, "preferred": true, "summary": {"categories": [["Carrier", 4454], ["Hosted", 2209], ["Unknown", 128]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "000008": {"column_number": 8, "datatype": "string", "name": "phone_model", "optype": "categorical", "order": 8, "preferred": true, "summary": {"categories": [["iPhone", 2834], ["Unknown", 1531], ["Android", 1470], ["Other", 550], ["Blackberry", 200], ["iPad", 60], ["Windows", 33], ["unknown", 10], ["Palm", 9], ["iphone", 8], ["android", 5], ["BlackBerry", 2], ["Iphone", 2], ["Android verizon youmail", 1], ["Samsung Galaxy S3", 1], ["Windows Phone", 1], ["iPod", 1]], "missing_count": 73}, "term_analysis": {"enabled": true}}, "000009": {"column_number": 9, "datatype": "string", "name": "successful_test_calls", "optype": "categorical", "order": 9, "preferred": true, "summary": {"categories": [["1", 4954], ["0", 1837]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "00000a": {"column_number": 10, "datatype": "string", "name": "failed_test_calls", "optype": "categorical", "order": 10, "preferred": true, "summary": {"categories": [["0", 3609], ["1", 3182]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "00000b": {"column_number": 11, "datatype": "string", "name": "signup_source", "optype": "categorical", "order": 11, "preferred": true, "summary": {"categories": [["Website", 4030], ["Al", 1282], ["API", 595], ["Al iOS AppStore", 294], ["Al Android Play Store", 286], ["Mobile Web Signup", 276], ["Al iOS Cydia", 14], ["Blackberry Developers", 1]], "missing_count": 13}, "term_analysis": {"enabled": true}}, "00000c": {"column_number": 12, "datatype": "string", "name": "carrier_name", "optype": "categorical", "order": 12, "preferred": true, "summary": {"categories": [["AT&T Wireless", 3102], ["Verizon Wireless", 1719], ["Sprint", 883], ["T-Mobile", 674], ["Metro PCS", 148], ["Home Phone", 141], ["US Cellular", 46], ["Office Phone", 32], ["Fido", 14], ["Bell", 12], ["Rogers", 10], ["Telus", 6], ["Nextel", 3], ["Simple Mobile", 1]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "00000d": {"column_number": 13, "datatype": "string", "name": "activation_status", "optype": "categorical", "order": 13, "preferred": true, "summary": {"categories": [["1", 5651], ["0", 1140]], "missing_count": 0}, "term_analysis": {"enabled": true}}, "00000e": {"column_number": 14, "datatype": "int16", "name": "total_voicemail", "optype": "numeric", "order": 14, "preferred": true, "summary": {"bins": [[2.14376, 2678], [17.51028, 1654], [33.39724, 652], [47.87113, 613], [62.39844, 256], [78.375, 376], [97.39877, 163], [115.36082, 97], [128.61818, 55], [144.90667, 75], [161.95833, 24], [177.14286, 42], [200.11538, 26], [220.41667, 12], [233.83333, 6], [248.18182, 11], [268.75, 8], [285.66667, 6], [306.25, 4], [326.5, 2], [350.66667, 9], [371.42857, 7], [393, 1], [415, 3], [430, 2], [445, 1], [494, 2], [561, 2], [587, 1], [767, 1], [859, 1], [967, 1]], "maximum": 967, "mean": 32.29465, "median": 15.48809, "minimum": 0, "missing_count": 0, "population": 6791, "splits": [0.5474, 4.16766, 11.51628, 21.56771, 36.90372, 66.61375], "standard_deviation": 50.89354, "sum": 219313, "sum_squares": 24669773, "variance": 2590.15249}}, "00000f": {"column_number": 15, "datatype": "string", "name": "status", "optype": "categorical", "order": 15, "preferred": true, "summary": {"categories": [["active", 5634], ["canceled", 1157]], "missing_count": 0}, "term_analysis": {"enabled": true}}}, "fields_meta": {"count": 16, "limit": 1000, "offset": 0, "query_total": 16, "total": 16}, "input_fields": ["000000", "000001", "000002", "000003", "000004", "000005", "000006", "000007", "000008", "000009", "00000a", "00000b", "00000c", "00000d", "00000e", "00000f"], "locale": "en_US", "missing_numeric_rows": 0, "missing_tokens": ["", "NaN", "NULL", "N/A", "null", "-", "#REF!", "#VALUE!", "?", "#NULL!", "#NUM!", "#DIV/0", "n/a", "#NAME?", "NIL", "nil", "na", "#N/A", "NA"], "name": "ultimate_7k's dataset", "number_of_batchcentroids": 0, "number_of_batchpredictions": 0, "number_of_clusters": 0, "number_of_ensembles": 0, "number_of_evaluations": 0, "number_of_models": 1, "number_of_predictions": 502, "objective_field": {"column_number": 15, "datatype": "string", "id": "00000f", "name": "status", "optype": "categorical", "order": 15, "term_analysis": {"enabled": true}}, "price": 0.0, "private": true, "ranges": null, "replacements": null, "resource": "dataset/53e3dd0fffa0440f4300a95c", "rows": 6791, "sample_rates": null, "seeds": null, "shared": false, "size": 499443, "source": "source/53e3dd010af5e83906003c02", "source_status": true, "status": {"bytes": 499443, "code": 5, "elapsed": 1741, "field_errors": [], "message": "The dataset has been created", "progress": 1.0, "row_format_errors": [], "serialized_rows": 6791, "task": "Done"}, "subscription": false, "tags": [], "term_limit": 32, "updated": "2014-08-07T20:21:21.440000", "user_metadata": {}}`)}},
			},
		}

		dataset, err := bigml.GetDataset("53e3dd0fffa0440f4300a95c")
		So(err, ShouldBeNil)
		So(dataset, ShouldNotBeNil)

		So(dataset.AllFields, ShouldEqual, true)
		So(dataset.Resource, ShouldEqual, "dataset/53e3dd0fffa0440f4300a95c")
	})
}
