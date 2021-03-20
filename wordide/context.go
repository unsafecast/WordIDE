package wordide

import (
	"archive/zip"
	"fmt"
)

type Context struct {
	Reader *zip.ReadCloser
}

func OpenContext(filePath string) (*Context, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}

	return &Context{Reader: r}, nil
}

func (ctx *Context) GetFile(path string) (*zip.File, error) {
	for _, f := range ctx.Reader.File {
		if f.Name == path {
			return f, nil
		}
	}

	return nil, fmt.Errorf("could not find file %s in context", path)
}

func (ctx *Context) Close() {
	ctx.Reader.Close()
}
