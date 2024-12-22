package text

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "positive",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Parser{
				root:        "testdata",
				UserService: nil,
				rows:        make(chan string, chanBuffer),
				filePaths:   make(chan string, chanBuffer),
			}

			err := p.Parse(ctx)
			tt.wantErr(t, err)
		})
	}
}
