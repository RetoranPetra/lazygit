package tag

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Merge = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Checkout a tag",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("test.txt","1")
		shell.Commit("first")
		shell.NewBranch("tagbranch")
		shell.UpdateFile("test.txt","2")
		shell.GitAddAll()
		shell.Commit("tagcommit")
		shell.CreateLightweightTag("tag", "HEAD")
		shell.Checkout("master")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Tags().
			Focus().
			Lines(
				Contains("tag").IsSelected(),
			).Press(keys.Branches.MergeIntoCurrentBranch)
	},
})
