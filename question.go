package collecta

import (
	"crypto/md5"
	"encoding/hex"
)

type Question struct {
	Order     int
	ID        string
	Title     string
	Anonymous bool
	Input     Input
}

func (q *Question) CalculateID() {
	hash := md5.Sum([]byte(q.Title))
	q.ID = hex.EncodeToString(hash[:])
}
