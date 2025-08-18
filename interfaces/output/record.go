package output

type RecordResponse struct {
	Data      interface{}
	page      int
	Total     int
	totalPage int
}
