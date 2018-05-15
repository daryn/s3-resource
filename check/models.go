package check

import "github.com/daryn/s3-resource"

type Request struct {
	Source  s3resource.Source  `json:"source"`
	Version s3resource.Version `json:"version"`
}

type Response []s3resource.Version
