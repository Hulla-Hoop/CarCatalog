package carcatalog

import (
	repo "carcatalog/internal/REPO"
	"carcatalog/internal/config"
	"carcatalog/internal/model"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_carcatalog_checkcar(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repo
		cfg    *config.ConfigRemoteApi
	}
	type args struct {
		reqId string
		car   *model.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "one",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "1",
				car: &model.Car{
					Id:   1,
					Mark: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "two",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "1",
				car: &model.Car{
					Id:   1,
					Mark: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carcatalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			if err := c.checkcar(tt.args.reqId, tt.args.car); (err != nil) != tt.wantErr {
				t.Errorf("carcatalog.checkcar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
