// Back-End in Go server
package psql

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
)

// go test -v -run ^TestDbConnect$
func TestDbConnect(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test_db_connect_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got := NewConn(&ctx, ConfigLog)

			if err := got.Conn.Ping(*got.Ctx); err != nil {
				t.Errorf("DbConnect() error = %v, got =%v", err, got)
				return
			}
		})
	}
}
