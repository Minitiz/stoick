package db

type Data struct {
	URL    string `json:"url"`
	Short  string `json:"short"`
	Access int64  `json:"access"`
	Expiry int64  `json:"expiry"`
}
