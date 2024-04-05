package psql

import (
	"carcatalog/internal/model"
	"database/sql"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_psql_getQuery(t *testing.T) {
	type fields struct {
		dB     *sql.DB
		logger *logrus.Logger
	}
	type args struct {
		reqId string
		carDB *model.CarDB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "one",
			fields: fields{
				dB:     nil,
				logger: logrus.New(),
			},
			args: args{
				reqId: "1",
				carDB: &model.CarDB{
					Id:   1,
					Mark: "test",
				},
			},
			want: "UPDATE cars SET mark = 'test', updated_at = NOW() WHERE id = 1 returning *;",
		},
		{
			name: "two",
			fields: fields{
				dB:     nil,
				logger: logrus.New(),
			},
			args: args{
				reqId: "2",
				carDB: &model.CarDB{
					Id:   1,
					Mark: "test",
					Year: 2020,
					Name: "test",
				},
			},
			want: "UPDATE cars SET mark = 'test', year = 2020, name = 'test', updated_at = NOW() WHERE id = 1 returning *;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &psql{
				dB:     tt.fields.dB,
				logger: tt.fields.logger,
			}
			if got := p.getQuery(tt.args.reqId, tt.args.carDB); strings.EqualFold(got, tt.want) {
				t.Errorf("psql.getQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
