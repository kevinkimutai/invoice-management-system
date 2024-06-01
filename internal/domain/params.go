package domain

type Params struct {
	Page  string
	Limit string
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

func CheckParams(m map[string]string) Params {

	params := Params{}

	if m["page"] != "" {
		params.Page = m["page"]
	}

	if m["limit"] != "" {
		params.Limit = m["limit"]
	}

	return params

}
