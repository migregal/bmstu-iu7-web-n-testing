.PHONY: load-test
load-test: HOST :=
load-test:
	k6 run docs-script.js --env HOST="$(HOST)"
	k6 run registration-script.js --env HOST="$(HOST)"
	k6 run model-uploading-script.js --env HOST="$(HOST)"
