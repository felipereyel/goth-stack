package embeded

import "embed"

// Embed a directory
//
//go:embed assets/*
var Assets embed.FS
