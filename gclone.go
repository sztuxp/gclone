package main

import (
	_ "github.com/going/gclone/backend/all" // import all backends
	_ "github.com/going/gclone/cmd/copy"
	"github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/all" // import all commands
	"github.com/rclone/rclone/fs"
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	fs.Version = fs.Version + "-mod1.3.1"
	cmd.Main()
}
