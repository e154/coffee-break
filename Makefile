DIR=build
PKG_ROOT=opt/watcher
PKG_NAME=watcher
VERSION=1.0.0

all: debian

debian:
	cd core/capi/src && $(MAKE) all
	go build
	npm install
	cd static_source && bower install
	gulp pack
	mkdir -p $(DIR)/$(PKG_ROOT)
	mkdir -p $(DIR)/$(PKG_ROOT)/lib
	mkdir -p $(DIR)/$(PKG_ROOT)/static_source
	cp watcher $(DIR)/$(PKG_ROOT)/watcher
	cp -r static_source/templates $(DIR)/$(PKG_ROOT)/static_source/templates
	cp -r static_source/js $(DIR)/$(PKG_ROOT)/static_source/js
	cp -r static_source/css $(DIR)/$(PKG_ROOT)/static_source/css
	cp -r static_source/audio $(DIR)/$(PKG_ROOT)/static_source/audio
	cp -r static_source/images $(DIR)/$(PKG_ROOT)/static_source/images
	./shared_library.sh watcher $(DIR)/$(PKG_ROOT)/lib
	cp watcher.sh $(DIR)/$(PKG_ROOT)/watcher.sh

	rm -rf $(DIR)/DEBIAN
	cp -r pkg/DEBIAN $(DIR)/DEBIAN
	cd $(DIR) && md5deep -r . > DEBIAN/md5sums
	fakeroot dpkg-deb --build build
	mv build.deb $(PKG_NAME)_$(VERSION).deb

clean:
	rm -f watcher
	rm -rf $(DIR)
	rm -rf node_modules
	rm -rf static_source/bower_components
	rm -rf static_source/css
	rm -rf static_source/js
	rm -rf static_source/node_modules
	cd core/capi/src && $(MAKE) clean

