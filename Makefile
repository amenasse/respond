.PHONY: build-respond
build-respond:
	${MAKE} -C cmd/respond build

install-respond:
	${MAKE} -C cmd/respond install
