package model

type SeearchResult struct {
	Hits  int64
	Start int
	Item  []interface{}
}
