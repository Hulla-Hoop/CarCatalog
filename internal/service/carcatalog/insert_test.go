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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &carcatalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			if got := c.check(tt.args.regNum); got != tt.want {
				t.Errorf("carcatalog.check() = %v, want %v", got, tt.want)
			}
		})
	}
}
