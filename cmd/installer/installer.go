package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/installer"
	"os"
)

var version string // ビルドスクリプトで埋め込む
var (
	app = kingpin.New("installer", "GoogleカレンダーTV会議強制起動ツールのインストラー")

	installCommand = app.Command("install", "インストール")

	updateCommand = app.Command("update", "アップデート")

	uninstallCommand = app.Command("uninstall", "アンインストール")
)

func main() {
	app.Version(version)
	appDir := common.GetAppDir()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case installCommand.FullCommand():
		inst := installer.NewInstaller()
		config := inst.ScanInput()
		// インストールする
		inst.Install(config, appDir)

	case updateCommand.FullCommand():
		if !common.FileExists(appDir) {
			fmt.Printf("インストールしたディレクトリが見つかりません. 探したパス: %s\n", appDir)
			os.Exit(1)
		}
		installer.NewUpdator().Update(appDir)
	case uninstallCommand.FullCommand():
		if !common.FileExists(appDir) {
			fmt.Printf("インストールしたディレクトリが見つかりません. 探したパス: %s\n", appDir)
			os.Exit(1)
		}
		installer.NewUninstaller().Uninstall(appDir)
	}
}
