# Copyright (C) 2021 Yu.

VERBOSE = no
CASE = ""

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  format  to format the code"
	@echo "  vet     to run golang vet"
	@echo "  lint    to run the staticcheck"
	@echo "  check   to format, vet, lint"
	@echo "  test    to run test case"
	@echo "  bench   to run benchmark test case"
	@exit 0

.PHONY: format
format:
	@[[ ${VERBOSE} = "yes" ]] && set -x; go fmt ./xgo/...;

.PHONY: vet
vet:
	@[[ ${VERBOSE} = "yes" ]] && set -x; go vet ./xgo/...;

.PHONY: lint
lint:
	@[[ ${VERBOSE} = "yes" ]] && set -x; staticcheck ./xgo/...;

.PHONY: check
check: format vet lint

#.PHONY: bench
#bench:
#	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -test.bench="." -test.run="Benchmark" -benchmem -count=1 ./...;

.PHONY: clean
clean:
	@[[ ${VERBOSE} = "yes" ]] && set -x; /bir/rm -fr ./build


.PHONY: generate-java
generate-java:
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_java.sh || bash scripts/generate_java.sh;

.PHONY: generate-go
generate-go:
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_go.sh || bash scripts/generate_go.sh;

.PHONY: generate-py
generate-py:
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_py.sh || bash scripts/generate_py.sh;

.PHONY: generate
generate: generate-go generate-java generate-py

.PHONY: compile
compile: generate-go
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/compile.sh || bash scripts/compile.sh;

.PHONY: install
install: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/install.sh || bash scripts/install.sh;

.PHONY: generate-test
generate-test: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/generate_test.sh || bash scripts/generate_test.sh;

.PHONY: test
test: generate-test check
	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -v -test.count=1 -failfast -test.run="${CASE}" ./xgo/tests;

.PHONY: test-only
test-only: generate-test
	@[[ ${VERBOSE} = "yes" ]] && set -x; go test -v -test.count=1 -failfast -test.run="${CASE}" ./xgo/tests;

.PHONY: bench
bench: generate-test vet
	@[[ ${VERBOSE} = "yes" ]] && set -x; cd tests; go test -test.bench="." -test.run="Benchmark" -benchmem -count=1 ./;

.PHONY: bench-only
bench-only:
	@[[ ${VERBOSE} = "yes" ]] && set -x; cd tests; go test -test.bench="." -test.run="Benchmark" -benchmem -count=1 ./;

.PHONY: test-json-error
test-json-error: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/test_gojson_error.sh || bash scripts/test_gojson_error.sh

.PHONY: test-defaults-error
test-defaults-error: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/test_godefaults_error.sh || bash scripts/test_godefaults_error.sh

.PHONY: test-validator-error
test-validator-error: compile
	@[[ ${VERBOSE} = "yes" ]] && bash -x scripts/test_govalidator_error.sh || bash scripts/test_govalidator_error.sh

.PHONY: test-error
test-error: test-json-error test-defaults-error test-validator-error

.DEFAULT_GOAL = help

# Target name % means that it is a rule that matches anything, @: is a recipe;
# the : means do nothing
%:
	@:

