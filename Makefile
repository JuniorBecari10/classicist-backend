BACK=back
SETUP=setup

back_build:
	cd $(BACK) && \
		make build

back_run:
	cd $(BACK) && \
		make run

back_watch:
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

build: back_build setup_build
