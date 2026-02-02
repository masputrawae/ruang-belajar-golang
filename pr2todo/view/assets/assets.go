package assets

import "embed"

//go:embed favicon.ico
var Favicon embed.FS

//go:embed main.css
var Css embed.FS

//go:embed main.js
var Javascript embed.FS
