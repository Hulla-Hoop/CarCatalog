package carcatalog

import (
	repo "carcatalog/internal/REPO"
	"carcatalog/internal/config"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_carcatalog_idCheckAndConvert(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repo
		cfg    *config.ConfigRemoteApi
	}
	type args struct {
		reqId string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carcatalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			got, err := c.idCheckAndConvert(tt.args.reqId, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("carcatalog.idCheckAndConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("carcatalog.idCheckAndConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
