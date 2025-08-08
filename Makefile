include pkg/testdata/Makefile

test_dir=pkg/testdata

install:
	go install

test:
	cd $(test_dir) && make test-package
