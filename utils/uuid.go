package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

func UUID() string {
	if id, err := uuid.NewV4(); nil == err {
		return id.String()
	} else {
		rand.Seed(time.Now().UnixNano())
		return fmt.Sprintf("%v%v", rand.Int63(), time.Now().UnixNano())
	}
}
