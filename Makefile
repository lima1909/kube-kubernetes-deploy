.PHONY: test
test:
	go test -count=1  ./...


.PHONY: deps
deps:
	govendor add +external github.com/openshift/origin/^
	govendor add +external github.com/openshift/client-go/^


# more documentation
pre-dep:
	go get -u -d github.com/openshift/api
	git checkout tags/v3.9.0