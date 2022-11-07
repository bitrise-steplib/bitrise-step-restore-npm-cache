package step

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-steputils/v2/cache"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
)

const stepId = "restore-npm-cache"

// Cache key templates
// OS + Arch: guarantees unique cache per stack (node_modules can contain compiled binaries)
// checksum: NPM or Yarn lockfiles
var keys = []string{
	`{{ .OS }}-{{ .Arch }}-npm-cache-{{ checksum "package-lock.json" "yarn.lock" }}`,
	`{{ .OS }}-{{ .Arch }}-npm-cache-`,
}

type Input struct {
	Verbose bool `env:"verbose,required"`
}

type RestoreCacheStep struct {
	logger      log.Logger
	inputParser stepconf.InputParser
	envRepo     env.Repository
	cmdFactory  command.Factory
}

func New(
	logger log.Logger,
	inputParser stepconf.InputParser,
	envRepo env.Repository,
	cmdFactory command.Factory,
) RestoreCacheStep {
	return RestoreCacheStep{
		logger:      logger,
		inputParser: inputParser,
		envRepo:     envRepo,
		cmdFactory:  cmdFactory,
	}
}

func (step RestoreCacheStep) Run() error {
	var input Input
	if err := step.inputParser.Parse(&input); err != nil {
		return fmt.Errorf("failed to parse inputs: %w", err)
	}
	stepconf.Print(input)
	step.logger.Println()
	step.logger.Printf("Cache keys:")
	step.logger.Printf(strings.Join(keys, "\n"))
	step.logger.Println()

	step.logger.EnableDebugLog(input.Verbose)

	restorer := cache.NewRestorer(step.envRepo, step.logger, step.cmdFactory)
	return restorer.Restore(cache.RestoreCacheInput{
		StepId:  stepId,
		Verbose: input.Verbose,
		Keys:    keys,
	})
}
