package data

import "embed"

// This package embeds data for the application.
//
// For deployment, distribute these directories alongside the binary:
// ├── finch-game				 # Binary
// └── data/
//     ├── assets/				 # Runtime assets (textures, levels, etc.)
//     └── resource.manifest	 # Metadata for resource routing
//
// The embedded/ directory is embedded into the binary and does not need to be distributed.

//go:embed embedded/*
var EmbeddedFS embed.FS
