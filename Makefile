clean-coverage:
	find . -name "coverage.out" -type f -delete

run-coverage:
	make clean-coverage
	./run-coverage.sh