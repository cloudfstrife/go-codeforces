PROJECT_NAME = go-codeforces

# main version
VERSION ?= $(shell git describe --tags --always --dirty)

# git commit Hash
COMMIT_HASH ?= $(shell git show -s --format=%H)

# build at
BUILD_TIME ?= $(shell date +"%F %T")

# go source file list
GOFILES := $(shell find . ! -path "./vendor/*" -name "*.go")

# unit test environment
TEST_ENV := 

# benchmark test environment
BENCHMARK_ENV := 

# unit test options
TEST_OPTS := 

# benchmark test options
BENCHMARK_OPTS := -cpu 1,2,3,4,5,6,7,8 -benchmem 

# sonar report output folder
REPORT_FOLDER := sonar

# sonar report file list
TEST_REPORT := ${REPORT_FOLDER}/test.report 
COVER_REPORT := ${REPORT_FOLDER}/cover.report
GOLANGCI_LINT_REPORT := ${REPORT_FOLDER}/golangci-lint.xml 
GOLINT_REPORT := ${REPORT_FOLDER}/golint.report 
GO_VET_REPORT := ${REPORT_FOLDER}/vet.report 

SONAR_SERVICE=http://xxx.xxx.xxx.xxx:9000
SONAR_TOKEN=

.PHONY: format test benchmark sonar clean
.DEFAULT_GOAL := test

# format go code
format:
	@for f in ${GOFILES} ; do 															\
		gofmt -w $${f};																	\
	done

# unit test
test: 
	${TEST_ENV} go test ${TEST_OPTS} ./...

# benchmark test
benchmark:
	${BENCHMARK_ENV} go test -bench . -run ^$$ ${BENCHMARK_OPTS}  ./...

# sonar
sonar: 
	mkdir -p ${REPORT_FOLDER}
	go test -json ./... > ${TEST_REPORT}
	go test -coverprofile=${COVER_REPORT} ./... 
	golangci-lint run --out-format checkstyle --issues-exit-code 0 ./... > ${GOLANGCI_LINT_REPORT}
	go vet ./... > ${GO_VET_REPORT} 2>&1 
	sonar-scanner -Dsonar.projectKey=${PROJECT_NAME} \
	-Dsonar.sources=. \
	-Dsonar.host.url=${SONAR_SERVICE} \
	-Dsonar.login=${SONAR_TOKEN} \
	-Dproject.settings=sonar/sonar-project.properties


# clean target executable program and sonar report
clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${COVER_REPORT}
	-rm -f ${GOLANGCI_LINT_REPORT}
	-rm -f ${GOLINT_REPORT}
	-rm -f ${GO_VET_REPORT}
	-go clean 
	-go clean -cache