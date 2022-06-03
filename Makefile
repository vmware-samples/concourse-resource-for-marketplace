.PHONY: lint test set-pipeline

test-lib/shunit2/shunit2:
	vendir sync

TEST_FILES = $(shell ls test/*.sh)
lint: bin/check bin/in bin/out $(TEST_FILES)
	shellcheck $?

test: bin/check bin/in bin/out $(TEST_FILES) test-lib/shunit2/shunit2
	ls $(TEST_FILES) | xargs -t -n 1 bash -c

set-pipeline: ci/pipeline.yaml
	fly -t tie set-pipeline --config ci/pipeline.yaml --pipeline marketplace-concourse-resource
