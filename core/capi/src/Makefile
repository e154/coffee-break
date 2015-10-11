CC=g++
CFLAGS=-c -Wall -std=c++1y -O0 -g3 -fmessage-length=0 -fPIC
LDFLAGS=
SOURCES=capi.cpp mainwindow.cpp moc_systemtray.cpp systemtray.cpp
LIBS= -lQt5Widgets -lQt5WebKitWidgets -lQt5WebKit -lQt5Core -lQt5Gui -lstdc++
#INCLUDES= -I/usr/include/i386-linux-gnu/qt5 -I/usr/include/i386-linux-gnu/qt5/QtGui -I/usr/include/i386-linux-gnu/qt5/QtWidgets -I/usr/include/i386-linux-gnu/qt5/QtWebKit -I/usr/include/i386-linux-gnu/qt5/QtWebKitWidgets -I/usr/include/i386-linux-gnu/qt5/QtCore
INCLUDES= -I/usr/include/x86_64-linux-gnu/qt5 -I/usr/include/x86_64-linux-gnu/qt5/QtGui -I/usr/include/x86_64-linux-gnu/qt5/QtWidgets -I/usr/include/x86_64-linux-gnu/qt5/QtWebKit -I/usr/include/x86_64-linux-gnu/qt5/QtWebKitWidgets -I/usr/include/x86_64-linux-gnu/qt5/QtCore
OBJECTS=$(SOURCES:.cpp=.o)
SHAREDLIB=libcapi.so
STATICLIB=libcapi.a

all: moc libcapi

moc:
	./update-moc.sh

libcapi: capi.o mainwindow.o moc_systemtray.o systemtray.o
#	$(CC) capi.o mainwindow.o moc_systemtray.o systemtray.o -shared -o $(SHAREDLIB) $(LIBS) -L/usr/lib/x86_64-linux-gnu/qt5
	ar -r  ../$(STATICLIB) ./capi.o ./mainwindow.o ./moc_systemtray.o ./systemtray.o

.SUFFIXES: .cpp .o

.cpp.o:
	$(CC) $(CFLAGS) $(INCLUDES) -c -o $@ $<

clean:
	rm -rf *.o libcapi
