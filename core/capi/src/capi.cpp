/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   capi.cpp
 * Version:     v1.00
 * Created:     23:42:52 / 18 авг. 2015 г.
 * Author:      delta54 <support@e154.ru>
 * 
 * Your use and or redistribution of this software in source and / or
 * binary form, with or without modification, is subject to: (i) your
 * ongoing acceptance of and compliance with the terms and conditions of
 * the DeltaSync License Agreement; and (ii) your inclusion of this notice
 * in any version of this software that you use or redistribute.
 * A copy of the DeltaSync License Agreement is available by contacting
 * DeltaSync Studios. at http://www.inet-print.ru/
 *
 * Description: 
 * ------------------------------------------------------------------
 * History:
 *
 */


#include <iostream>

#include <QApplication>
#include <QDebug>
#include <QAction>
#include <QRadioButton>
#include <QObject>
#include <QWidget>
#include <QDialog>
#include <QMenu>

#include "systemtray.h"
#include "capi.h"

// Application
void NewGuiApplication() {
    static char empty[1] = {0};
    static char *argv[] = {empty, 0};
    static int argc = 1;
    new QApplication(argc, argv);

    // The event loop should never die.
    qApp->setQuitOnLastWindowClosed(false);
}

void ApplicationExec() { qApp->exec(); }
void ApplicationExit() { qApp->exit(0); }
void ApplicationFlushAll() { qApp->processEvents(); }
QApp_ *ApplicationPtr() { return qApp; };

// singleton
// ----------------------------------------------------------------------------
SystemTray_ *GetSystemTray() {
	return SystemTray::getSingletonPtr();
}

void SetTrayIcon(SystemTray_ *t, char *img) {
	reinterpret_cast<SystemTray *>(t)->setTrayIcon(img);
}

void SetTrayToolTip(SystemTray_ *t, char *tooltip) {
	reinterpret_cast<SystemTray *>(t)->setTrayToolTip(tooltip);
}

void SetTrayVisible(SystemTray_*t, bool trigger) {
	reinterpret_cast<SystemTray *>(t)->setTrayVisible(trigger);
}

void ShowMessage(SystemTray_*t, char *title, char *msg, int icon) {
	reinterpret_cast<SystemTray *>(t)->showMessage(QString(title), QString(msg), QSystemTrayIcon::MessageIcon(icon));
}

// time
// ----------------------------------------------------------------------------
void SetTime(SystemTray_ *t, int time) {
	reinterpret_cast<SystemTray *>(t)->setTimer(time, 0);
}

int GetTime(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getTime();
}

void SetTimeCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setTimeCallback(callback);
}


