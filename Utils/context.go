package utils 

import (
	"context"
	"time"
)

func GetContext (timeout time.Duration)(context.Context,context.CancelFunc) {
     return context.WithTimeout(context.Background(),timeout)
}