package collecta

import (
	"crypto/md5"
	"encoding/hex"
)

type Question struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Anonymous bool   `json:"anonymous"`
	Input     Input  `json:"input"`
}

func (q *Question) calculateID() {
	hash := md5.Sum([]byte(q.Title))
	q.ID = hex.EncodeToString(hash[:])
}
