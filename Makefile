DIR=build
PKG_ROOT=opt/coffeebreak
PKG_NAME=coffeebreak
VERSION=1.3.4

all: debian

debian:
	cd core/capi/src && $(MAKE) all
	go build -ldflags="-X settings.APP_VERSION='${VERSION}'"
	npm install
	cd static_source && bower install
	gulp pack
	mkdir -p $(DIR)/$(PKG_ROOT)
	mkdir -p $(DIR)/usr
	mkdir -p $(DIR)/$(PKG_ROOT)/static_source
	cp coffee-break $(DIR)/$(PKG_ROOT)/coffee-break
	cp -r static_source/templates $(DIR)/$(PKG_ROOT)/static_source/templates
	cp -r static_source/js $(DIR)/$(PKG_ROOT)/static_source/js
	cp -r static_source/css $(DIR)/$(PKG_ROOT)/static_source/css
	cp -r static_source/audio $(DIR)/$(PKG_ROOT)/static_source/audio
	cp -r static_source/images $(DIR)/$(PKG_ROOT)/static_source/images
	cp pkg/share $(DIR)/usr/share -r
	cp pkg/coffee-break.sh $(DIR)/$(PKG_ROOT)/coffee-break.sh
	chmod +x $(DIR)/$(PKG_ROOT)/coffee-break.sh
	rm -rf $(DIR)/DEBIAN
	cp -r pkg/DEBIAN $(DIR)/DEBIAN

	# dynamic app pack
	# set package version
	sed 's/Version: \(.*\)/Version: ${VERSION}/g' $(DIR)/DEBIAN/control > $(DIR)/DEBIAN/control_t
	mv $(DIR)/DEBIAN/control_t $(DIR)/DEBIAN/control
	cd $(DIR) && md5deep -r . > DEBIAN/md5sums
	fakeroot dpkg-deb --build build
	mv build.deb $(PKG_NAME)_$(VERSION)_amd64.deb

	# static app pack
	sed 's/Depends: \(.*\)/Depends: /g' $(DIR)/DEBIAN/control > $(DIR)/DEBIAN/control_t
	mv $(DIR)/DEBIAN/control_t $(DIR)/DEBIAN/control
	sed 's/Package: \(.*\)/Package: coffee-break-static/g' $(DIR)/DEBIAN/control > $(DIR)/DEBIAN/control_t
	mv $(DIR)/DEBIAN/control_t $(DIR)/DEBIAN/control
	./shared_library.sh ${DIR}/${PKG_ROOT}/coffee-break ${DIR}/${PKG_ROOT}/lib
	cd ${DIR}/${PKG_ROOT}/lib

	# copy qt5 plugins
	mkdir -p ${DIR}/${PKG_ROOT}/platforms
	cp /usr/lib/x86_64-linux-gnu/qt5/plugins/platforms/libqxcb.so ${DIR}/${PKG_ROOT}/platforms/libqxcb.so
	./shared_library.sh ${DIR}/${PKG_ROOT}/platforms/libqxcb.so ${DIR}/${PKG_ROOT}/lib

	ls | grep -v "lib[Qt5 z xcb c pthread]" | rm -f
	cd ../../

	# create package
	fakeroot dpkg-deb --build build
	mv build.deb $(PKG_NAME)_$(VERSION)_amd64_static.deb

clean:
	rm -f ${PKG_NAME}
	rm -rf $(DIR)
	rm -rf node_modules
	rm -rf static_source/bower_components
	rm -rf static_source/css
	rm -rf static_source/js
	rm -rf static_source/node_modules
	cd core/capi/src && $(MAKE) clean
	rm -f *.deb

