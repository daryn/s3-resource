package in

import "github.com/daryn/s3-resource"

type Request struct {
	Source  s3resource.Source  `json:"source"`
	Version s3resource.Version `json:"version"`
	Params  Params             `json:"params"`
}

type Params struct {
	Unpack bool `json:"unpack"`
}

type Response struct {
	Version  s3resource.Version        `json:"version"`
	Metadata []s3resource.MetadataPair `json:"metadata"`
}
