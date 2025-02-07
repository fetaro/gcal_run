package gcal_run

import (
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	credentialPath := "/tmp/gcal_run_test/credential.json"
	if !common.FileExists(credentialPath) {
		t.Skip("クレデンシャルファイルが存在しないので、テストはスキップ")
	}
	config := common.NewConfig("/tmp/gcal_run_test/credential.json", 30, "/Applications/Google Chrome.app")
	runner := NewRunner(config, "/tmp/gcal_run_test")
	err := runner.Run()
	assert.NoError(t, err)
}
