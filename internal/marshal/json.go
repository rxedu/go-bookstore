package marshal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ParseID(ID string) (int64, bool) {
	val, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Printf("cannot parse id as int")
		return 0, false
	}
	return val, true
}
