package memory

import (
	"git.tz.com/devops/gin-core/config/loader"
	"git.tz.com/devops/gin-core/config/reader"
	"git.tz.com/devops/gin-core/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
