package commons

import "github.com/minskylab/collecta/api/graph/model"

// PtrBool only returns a boolean pointer instance of a boolean value
func PtrBool(v bool) *bool {
	return &v
}

func MapToPairs(m map[string]string) []*model.Pair {
	pairs := make([]*model.Pair, 0)
	for k, v := range m {
		pairs = append(pairs, &model.Pair{
			Key:   k,
			Value: v,
		})
	}
	return pairs
}