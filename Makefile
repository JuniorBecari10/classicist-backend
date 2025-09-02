BACK=backend
SETUP=setup

backend_build:
	cd $(BACK) && \
		make build

backend_run:
	cd $(BACK) && \
		make run

backend_watch:
	cd $(BACK) && \
		make watch

# ---

setup_build:
	cd $(SETUP) && \
		make build

setup_run:
	cd $(SETUP) && \
		make run

setup_watch:
	cd $(SETUP) && \
		make watch

# ---

build: backend_build setup_build
