.PHONY : all vet fmt

all : vet fmt test

vet :
	go vet github.com/tangzixiang/mplus
	go vet github.com/tangzixiang/mplus/context
	go vet github.com/tangzixiang/mplus/data
	go vet github.com/tangzixiang/mplus/decode
	go vet github.com/tangzixiang/mplus/errs
	go vet github.com/tangzixiang/mplus/header
	go vet github.com/tangzixiang/mplus/message
	go vet github.com/tangzixiang/mplus/mhttp
	go vet github.com/tangzixiang/mplus/middleware
	go vet github.com/tangzixiang/mplus/mime
	go vet github.com/tangzixiang/mplus/plus
	go vet github.com/tangzixiang/mplus/query
	go vet github.com/tangzixiang/mplus/route
	go vet github.com/tangzixiang/mplus/validate
	go vet github.com/tangzixiang/mplus/util

fmt :
	gofmt -l -e -d -w .

test : 
	go test