package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stmp), nil
}

type UserRespone struct {
	ID       int32  `json:"id"`
	NickName string `json:"name"`
	Gender   string `json:"gender"`
	// BirthDay string `json:"birthday"`
	BirthDay JsonTime `json:"birthday"`
	Mobile   string   `json:"mobile"`
}
