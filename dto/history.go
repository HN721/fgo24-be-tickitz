package dto

type HistoryReq struct {
	IdTransaction int    `json:"id_trx"`
	Status        string `json:"status"`
	Note          string `json:"note"`
}
