package postgres

// Config contains Postgres configuration
type Config struct {
	URL                  string `required:"true"`
	MaxConnections       int32  `split_words:"true" default:"10"`
	IdleConnections      int32  `split_words:"true" default:"1"`
	PreferSimpleProtocol bool   `split_words:"true"`
}
