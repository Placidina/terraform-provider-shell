package shell

//Config is the config for the client.
type Config struct {
	Arguments            []interface{}
	Environment          map[string]interface{}
	SensitiveEnvironment map[string]interface{}
	Interpreter          []interface{}
	EnableParallelism    bool
}

//Client is the client itself. Since we already have access to the shell no real provisioning needs to be done
type Client struct {
	config *Config
}

// Client configures and returns a fully initialized ShellClient
func (c *Config) Client() (*Client, error) {
	client := &Client{
		config: c,
	}

	return client, nil
}
