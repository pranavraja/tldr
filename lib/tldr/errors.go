package tldr

import "errors"

var ErrNotFound = errors.New("Not found.\nTo add this command, send a pull request at:\n  https://github.com/tldr-pages/tldr")
