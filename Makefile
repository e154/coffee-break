
all: watcher

watcher:
	cd core/capi/src && $(MAKE) all
	go build
	npm install
	cd static_source && bower install
	gulp pack

clean:
	rm -f watcher
	cd core/capi/src && $(MAKE) clean

