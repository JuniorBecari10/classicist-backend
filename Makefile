BE=backend

build:
	cd $(BE) && \
		make build

run:
	cd $(BE) && \
		make run

watch:
	cd $(BE) && \
		make watch
