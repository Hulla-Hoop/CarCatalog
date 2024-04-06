package carcatalog

import (
	repo "carcatalog/internal/REPO"
	"carcatalog/internal/config"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_carcatalog_checkFilter(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repo
		cfg    *config.ConfigRemoteApi
	}
	type args struct {
		reqId    string
		limit    string
		offset   string
		field    string
		value    string
		operator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name: "one",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId:    "1",
				limit:    "1",
				offset:   "1",
				field:    "mark",
				value:    "Ford",
				operator: "eq",
			},
			want: map[string]string{"limit": "1", "offset": "1", "field": "mark", "value": "Ford", "operator": "eq"},
		},
		{
			name: "two",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId:    "1",
				limit:    "",
				offset:   "",
				field:    "",
				value:    "",
				operator: "",
			},
			want: map[string]string{},
		},
		{
			name: "three",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId:    "1",
				limit:    "ad",
				offset:   "dds",
				field:    "ds",
				value:    "er",
				operator: "tar",
			},
			want: map[string]string{"value": "er"},
		},
		{
			name: "four",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId:    "1",
				limit:    "1",
				offset:   "1",
				field:    "",
				value:    "",
				operator: "",
			},
			want: map[string]string{"limit": "1", "offset": "1"},
		},
		{
			name: "five",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId:    "1",
				limit:    "-1",
				offset:   "-31",
				field:    "model",
				value:    "1",
				operator: "eq",
			},
			want: map[string]string{"field": "model", "value": "1", "operator": "eq"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carcatalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			if got := c.checkFilter(tt.args.reqId, tt.args.limit, tt.args.offset, tt.args.field, tt.args.value, tt.args.operator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("carcatalog.checkFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
