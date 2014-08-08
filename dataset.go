package bigml

type Meta struct {
	Limit      int     `json:"limit"`
	Next       *string `json:"next"`
	Offset     int     `json:"offset"`
	Previous   *string `json:"previous"`
	TotalCount int     `json:"total_count"`
}

type Field struct {
	ColumnNumber int    `json:"column_number"`
	Datatype     string `json:"datatype"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Optype       string `json:"optype"`
	Order        int    `json:"order"`
	TermAnalysis struct {
		Enabled bool `json:"enabled"`
	} `json:"term_analysis"`
}

type Dataset struct {
	AllFields                bool             `json:"all_fields"`
	Category                 int              `json:"category"`
	Code                     int              `json:"code"`
	Columns                  int              `json:"columns"`
	Created                  string           `json:"created"`
	Credits                  float64          `json:"credits"`
	Description              string           `json:"description"`
	Dev                      bool             `json:"dev"`
	FieldTypes               map[string]int   `json:"field_types"`
	Fields                   map[string]Field `json:"fields"`
	Locale                   string           `json:"locale"`
	MissingNumericRows       int              `json:"missing_numeric_rows"`
	Name                     string           `json:"name"`
	NumberOfBatchCentroids   int              `json:"number_of_batchcentroids"`
	NumberOfBatchPredictions int              `json:"number_of_batchpredictions"`
	NumberOfClusters         int              `json:"number_of_clusters"`
	NumberOfEnsembles        int              `json:"number_of_ensembles"`
	NumberOfEvaluations      int              `json:"number_of_evaluations"`
	NumberOfModels           int              `json:"number_of_models"`
	NumberOfPredictions      int              `json:"number_of_predictions"`
	ObjectiveField           Field            `json:"objective_field"`
	Price                    float64          `json:"price"`
	Private                  bool             `json:"private"`
	Resource                 string           `json:"resource"`
	Rows                     int              `json:"rows"`
	Shared                   bool             `json:"shared"`
	Size                     int              `json:"size"`
	Source                   string           `json:"source"`
	SourceStatus             bool             `json:"source_status"`
}

type DatasetResponse struct {
	Meta     Meta      `json:"meta"`
	Datasets []Dataset `json:"objects"`
}
