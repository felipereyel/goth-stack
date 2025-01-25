package embeded

import "embed"

// Embed a directory
//
//go:embed migrations/*.sql
var Migrations embed.FS
