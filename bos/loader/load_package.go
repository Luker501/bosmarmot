package loader

import (
	"fmt"
	"path/filepath"

	"github.com/monax/bosmarmot/bos/def"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func LoadPackage(fileName string) (*def.Package, error) {
	log.Info("Loading monax Jobs Definition File.")
	var pkg = new(def.Package)
	var epmJobs = viper.New()

	// setup file
	abs, err := filepath.Abs(fileName)
	if err != nil {
		return nil, fmt.Errorf("Sorry, the marmots were unable to find the absolute path to the monax jobs file.")
	}

	dir := filepath.Dir(abs)
	base := filepath.Base(abs)
	extName := filepath.Ext(base)
	bName := base[:len(base)-len(extName)]
	log.WithFields(log.Fields{
		"path": dir,
		"name": bName,
	}).Debug("Loading jobs file")

	epmJobs.SetConfigType("yaml")
	epmJobs.SetConfigName(bName)

	r, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// load file
	if err := epmJobs.ReadConfig(r); err != nil {
		return nil, fmt.Errorf("Sorry, the marmots were unable to load the monax jobs file. Please check your path: %v", err)
	}

	// marshall file
	if err := epmJobs.UnmarshalExact(pkg); err != nil {
		return nil, fmt.Errorf(`Sorry, the marmots could not figure that monax jobs file out.
			Please check that your epm.yaml is properly formatted: %v`, err)
	}

	// TODO more file sanity check (fail before running)

	return pkg, nil
}


