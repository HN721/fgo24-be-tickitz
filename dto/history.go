package dto

type HistoryReq struct {
	IdTransaction int    `json:"id_trx" db:"transaction_id"`
	Status        string `json:"status"`
	Note          string `json:"note"`
}
