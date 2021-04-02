package logos

import "embed"

//go:embed data/en/*
var EmbeddedBible embed.FS
