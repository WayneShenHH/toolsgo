package logsvc_test

import (
	"testing"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/logsvc"
)

func Test_log(t *testing.T) {
	repo := repositoryimpl.New()
	svc := logsvc.New(repo)
	svc.Read(0, 1)
}
