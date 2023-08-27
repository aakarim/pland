package config

import (
	"fmt"
	"io"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/gocode/gocodec"
)

type FileConfig struct {
	EditorCommand string `cue:"editorCommand" json:"editorCommand,omitempty"`
}

func ParseFile(r io.Reader) (*FileConfig, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("ReadAll(): %w", err)
	}
	ctx := cuecontext.New()
	g := gocodec.New((*cue.Runtime)(ctx), nil)
	val := ctx.CompileBytes(b)
	if err := val.Validate(); err != nil {
		return nil, fmt.Errorf("validating config file: %w", err)
	}
	var fc FileConfig
	if err := g.Encode(val, &fc); err != nil {
		return nil, fmt.Errorf("loading CUE into Go: %w", err)
	}
	return &fc, nil
}
