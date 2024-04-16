package tag

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Merge = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Merge a tag",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.NewBranch("test")
		shell.EmptyCommit("one")
		shell.EmptyCommit("two")
		shell.CreateLightweightTag("tag", "HEAD^")
		shell.Checkout("master")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Tags().
			Focus().
			Lines(
				Contains("tag").IsSelected(),
			).
			Press(keys.Branches.MergeIntoCurrentBranch) //Merge keybind

		t.Views().Branches().IsFocused().Lines(
			Contains("two").IsSelected(),
			Contains("master"),
		)
	},
})
