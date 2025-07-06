package dto

type HistoryReq struct {
	Id            int    `json:"id ,omitempty" db:"id"`
	IdTransaction int    `json:"id_trx" db:"transaction_id"`
	Status        string `json:"status"`
	Note          string `json:"note"`
}
