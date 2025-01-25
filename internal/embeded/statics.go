package embeded

import "embed"

// Embed a directory
//
//go:embed statics/*
var Statics embed.FS
