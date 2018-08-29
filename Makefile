.PHONY: test
test:
	go test -count=1  ./...


.PHONY: deps
deps:
	govendor add +external github.com/openshift/origin/^
	govendor add +external github.com/openshift/client-go/^