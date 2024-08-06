package celeritas

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}
	return nil
}
func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	c.Version = version // Used to get rid of not used warning, to be removed before long
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := c.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}