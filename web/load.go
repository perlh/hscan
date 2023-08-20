package web

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	xray_structs "hscan/web/pkg/xray/structs"
	"strings"
)

//go:embed "pocs/xray2/**"
var Pocs embed.FS

func LoadPoc(dir string, fileName string, Pocs embed.FS) (*xray_structs.Poc, error) {
	p := &xray_structs.Poc{}
	yamlFile, err := Pocs.ReadFile(dir + "/" + fileName)

	if err != nil {
		fmt.Printf("[-] load poc %s error1: %v\n", fileName, err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		fmt.Printf("[-] load poc %s error2: %v\n", fileName, err)
		return nil, err
	}
	return p, err
}

func initXaryPoc(xrayDir string) map[string]xray_structs.Poc {
	entries, err := Pocs.ReadDir(xrayDir)
	xrayPoc := make(map[string]xray_structs.Poc, 1)
	if err != nil {
		fmt.Printf("[-] init poc error: %v", err)
		return nil
	}
	for _, one := range entries {
		path := one.Name()

		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
			if poc, _ := LoadPoc(xrayDir, path, Pocs); poc != nil {
				if poc.Transport == "" {
					poc.Transport = "http"
				}
				xrayPoc[poc.Name] = *poc
			}
		}
	}
	return xrayPoc

}
