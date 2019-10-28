package types

type Block struct {
	Height     uint64 `json:"height"`
	Time       uint64 `json:"time"`
	Difficulty uint64 `json:"difficulty"`
	Hash       string `json:"hash"`
}
