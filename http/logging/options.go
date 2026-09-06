package logging

type Config struct {
    LogRequestBody  bool
    LogResponseBody bool
    IncludeHeaders  bool
    MaxBodySize     int
}

func defaultConfig() Config {
    return Config{
        LogRequestBody:  false,
        LogResponseBody: false,
        IncludeHeaders:  false,
        MaxBodySize:     4096,
    }
}

type Option func(*Config)

func WithRequestBody() Option {
    return func(c *Config) {
        c.LogRequestBody = true
    }
}

func WithResponseBody() Option {
    return func(c *Config) {
        c.LogResponseBody = true
    }
}

func WithHeaders() Option {
    return func(c *Config) {
        c.IncludeHeaders = true
    }
}

func WithMaxBodySize(size int) Option {
    return func(c *Config) {
        if size > 0 {
            c.MaxBodySize = size
        }
    }
}
