package config

import (
	"sync"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"log"
	"fmt"
	"errors"
)

var (
	defaultFilename  = "config.yml"
	c               *cfg
	mu              sync.Mutex
)

func Get() *cfg {
	if c == nil {
		mu.Lock()
		defer mu.Unlock()
		if c != nil {
			return c
		}
		// Try to load form ENV
		var filename string
		if envFilename := os.Getenv("APP_CONFIG"); envFilename != "" {
			filename = envFilename
		} else {
			// Default
			curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatalf("Can't get config path. Error: %s", err)
			}
			filename = fmt.Sprintf("%s/%s", curDir, defaultFilename)
		}
		if filename == "" {
			log.Fatal("Config filename is empty")
		}

		err := load(filename)
		if err != nil {
			log.Fatalf("Config file didn't load. Error: %s", err)
		}

	}
	return c
}

func load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	conf := cfg{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	c = &conf
	return nil
}

// After Reset(), config will be reloaded
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	c = nil
}

func CreateDefaultConfigFile() error {
	if exists(defaultFilename) {
		return errors.New("Config file already exists")
	}
	data, err := yaml.Marshal(Defaults())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(defaultFilename, data, 0666)
	if err != nil {
		return err
	}
	log.Printf("Config file created: %s", defaultFilename)
	return nil
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
