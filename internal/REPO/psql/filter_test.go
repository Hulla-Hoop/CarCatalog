package psql

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_psql_queryFilter(t *testing.T) {
	type fields struct {
		dB     *sql.DB
		logger *logrus.Logger
	}

	tests := []struct {
		name   string
		fields fields
		filter map[string]string
		want   string
	}{
		{
			name: "one",
			fields: fields{
				dB:     nil,
				logger: nil,
			},

			filter: map[string]string{"offset": "1"},

			want: "SELECT * FROM cars WHERE removed = false AND id > 1",
		},
		{
			name: "two",
			fields: fields{
				dB:     nil,
				logger: nil,
			},

			filter: map[string]string{"field": "name", "value": "test", "operator": "eq"},

			want: "SELECT * FROM cars WHERE removed = false AND name = 'test'",
		},
		{
			name: "three",
			fields: fields{
				dB:     nil,
				logger: nil,
			},

			filter: map[string]string{"offset": "2", "limit": "2"},

			want: "SELECT * FROM cars WHERE removed = false AND id > 2 LIMIT 2",
		},
		{
			name: "four",
			fields: fields{
				dB:     nil,
				logger: nil,
			},

			filter: map[string]string{"offset": "2", "limit": "2", "field": "name", "value": "test", "operator": "ge"},

			want: "SELECT * FROM cars WHERE removed = false AND name >= 'test' AND id > 2 LIMIT 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &psql{
				dB:     tt.fields.dB,
				logger: tt.fields.logger,
			}
			if got := p.queryFilter(tt.filter); strings.EqualFold(got, tt.want) {
				t.Errorf("psql.queryFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
