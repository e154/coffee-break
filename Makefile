DIR=build

all: watcher

watcher:
	cd core/capi/src && $(MAKE) all
	go build
	npm install
	cd static_source && bower install
	gulp pack

install:
	mkdir -p $(DIR)
	mkdir -p $(DIR)/static_source
	cp watcher $(DIR)/watcher
	cp -r static_source/templates $(DIR)/static_source/templates
	cp -r static_source/js build/static_source/js
	cp -r static_source/css build/static_source/css
	cp -r static_source/audio build/static_source/audio
	cp -r static_source/images build/static_source/images

clean:
	rm -f watcher
	rm -f build
	cd core/capi/src && $(MAKE) clean

