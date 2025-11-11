package cmd

import (
	"fmt"

	"github.com/dakshpareek/ctx/internal/display"
)

const advancedDescription = `
Advanced command. Most users should rely on the guided workflow:

  ctx ask     # sync + prompt
  ctx update  # mark skeletons current
  ctx bundle  # export full context

Use this command when you need lower-level control.`

func printAdvancedNotice(preferred string) {
	if preferred == "" {
		return
	}
	fmt.Println(display.Info("Tip: For everyday work, run '%s' for the guided experience.", preferred))
}
