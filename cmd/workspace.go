package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dakshpareek/ctx/internal/display"
	"github.com/dakshpareek/ctx/internal/fs"
	"github.com/dakshpareek/ctx/internal/types"
)

func ensureWorkspace(requireIndex bool) (string, string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", "", &types.Error{Code: types.ExitCodeFileSystem, Err: fmt.Errorf("determine working directory: %w", err)}
	}

	ctxDir := filepath.Join(wd, ctxDirName)
	if !fs.Exists(ctxDir) {
		fmt.Println(display.Warning("No .ctx/ workspace found here. Run 'ctx init' from your project root to bootstrap it."))
		return "", "", &types.Error{Code: types.ExitCodeUserError}
	}

	indexPath := filepath.Join(ctxDir, indexFileName)
	if requireIndex && !fs.Exists(indexPath) {
		fmt.Println(display.Warning("Your .ctx/ workspace is missing index.json. Run 'ctx sync' to rebuild it."))
		return "", "", &types.Error{Code: types.ExitCodeData}
	}

	return ctxDir, indexPath, nil
}
