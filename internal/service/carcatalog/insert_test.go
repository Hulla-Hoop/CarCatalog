package carcatalog

import (
	repo "carcatalog/internal/REPO"
	"carcatalog/internal/config"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_carcatalog_check(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repo
		cfg    *config.ConfigRemoteApi
	}
	type args struct {
		regNum string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "one",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				regNum: "test",
			},
			want: true,
		},
		{
			name: "two",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				regNum: "x120xx123",
			},
			want: false,
		},
		{
			name: "three",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				regNum: "x120xx1233",
			},
			want: true,
		},
		{
			name: "four",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				regNum: "x120xxx123",
			},
			want: true,
		},
		{
			name: "five",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				regNum: "xx120xx123",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carcatalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			if got := c.check(tt.args.regNum); got == tt.want {
				t.Errorf("carcatalog.check() = %v, want %v", got, tt.want)
			}
		})
	}
}
