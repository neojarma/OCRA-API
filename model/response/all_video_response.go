package response

import joins_model "ocra_server/model/joins"

type VideosResponse struct {
	Page   int                          `json:"page"`
	Limit  int                          `json:"limit"`
	Videos []*joins_model.HomeVideoJoin `json:"videos"`
}
