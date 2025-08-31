BE=backend
SE=setup

be_build:
	cd $(BE) && \
		make build

be_run:
	cd $(BE) && \
		make run

be_watch:
	cd $(BE) && \
		make watch

# ---

se_build:
	cd $(SE) && \
		make build

se_run:
	cd $(SE) && \
		make run

se_watch:
	cd $(SE) && \
		make watch

# ---

build: be_build se_build
