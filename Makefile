run-basic:
	cd basic && go build && ./basic

run-narrow-lane:
	cd narrow-lane && go build && ./narrow-lane

clean:
	cd basic && rm basic
