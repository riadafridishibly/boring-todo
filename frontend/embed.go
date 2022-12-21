package frontend

import "embed"

//go:generate yarn
//go:generate yarn build
//go:embed all:dist
var DistDir embed.FS
