package web

import (
	// "embed"

	"github.com/aceberg/LightAlert/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllHosts - all hosts
	AllHosts []models.Host
	// TemplHTML - embed templates
	//
	// //go:embed templates/*
	// TemplHTML embed.FS
)

// const TemplPath = "templates/"

// TemplPath - path to html templates
const TemplPath = "../../internal/web/templates/"

// Icon - favicon
const Icon = "AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAA9N/iAGtr/wA9H/8A+fn5AGvx/wBUtP8AejMAAL5SAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIiHdwAAAACIiId3dwAACIiIh3d3cACIiIiHd3d3AIiIIiMzN3cIiIiIh3d3d3iIiERBERd3eIiIiId3d3d4iIiIh3d3d3iIiERBERd3eIiIiId3d3dwiIhVVmZndwCIiIiHd3d3AAiIiId3d3AAAIiIh3d3AAAAAIiHdwAAD4HwAA4AcAAMADAACAAQAAgAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIABAACAAQAAwAMAAOAHAAD4HwAA"
