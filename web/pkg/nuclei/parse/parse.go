package parse

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolinit"
	"github.com/projectdiscovery/nuclei/v2/pkg/templates"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
	"go.uber.org/ratelimit"
	"hscan/utils/logger"
	"hscan/web/errors"
	"hscan/web/pkg/nuclei/structs"
)

var (
	ExecuterOptions protocols.ExecuterOptions
)

func InitExecuterOptions(rate int, timeout int) {
	fakeWriter := structs.FakeWrite{}
	progress := &structs.FakeProgress{}
	o := types.Options{
		RateLimit:               rate,
		BulkSize:                25,
		TemplateThreads:         25,
		HeadlessBulkSize:        10,
		HeadlessTemplateThreads: 10,
		Timeout:                 timeout,
		Retries:                 1,
		MaxHostError:            30,
	}
	err := protocolinit.Init(&o)
	if err != nil {
		logger.Error("Nuclei InitExecuterOptions error")
		return
	}

	catalog2 := catalog.New("")
	ExecuterOptions = protocols.ExecuterOptions{
		Output:      &fakeWriter,
		Options:     &o,
		Progress:    progress,
		Catalog:     catalog2,
		RateLimiter: ratelimit.New(rate),
	}

}

func ParsePoc(filename string) (*structs.Poc, error) {
	var err error
	//templates.Parse()
	poc, err := templates.Parse(filename, nil, ExecuterOptions)
	if err != nil {
		return nil, err
	}
	if poc == nil {
		return nil, nil
	}
	if poc.ID == "" {
		return nil, errors.New("Nuclei poc id can't be nil")
	}
	return poc, nil
}
