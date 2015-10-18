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
#include "mainwindow.h"
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
QThread_ *ApplicationThread() { return qApp->thread(); }

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

void MoveToThread(SystemTray_*t, QThread_ *thread) {
	reinterpret_cast<SystemTray *>(t)->moveToThread(reinterpret_cast<QThread *>(thread));
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

// default time
// ----------------------------------------------------------------------------
void SetDTime(SystemTray_ *t, int time) {
	reinterpret_cast<SystemTray *>(t)->setDTimer(time, 0);
}

int GetDTime(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getDtime();
}

void SetDTimeCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setDtimeCallback(callback);
}

// alarm
// ----------------------------------------------------------------------------
void SetAlarm(SystemTray_ *t, int state) {
	reinterpret_cast<SystemTray *>(t)->setAlarm(state, 0);
}

int GetAlarm(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getAlarm();
}

void SetAlarmCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setAlarmCallback(callback);
}

// run at startup
// ----------------------------------------------------------------------------
void SetRunAtStartup(SystemTray_ *t, int state) {
	reinterpret_cast<SystemTray *>(t)->setRunAtStartup(state);
}

int GetRunAtStartup(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getRunAtStartup();
}

void SetRunAtStartupCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setRunAtStartupCallback(callback);
}

// run at startup
// ----------------------------------------------------------------------------
void SetAlarmInfo(SystemTray_ *t, char *info) {
	reinterpret_cast<SystemTray *>(t)->setAlarmInfo(info);
}

char *GetAlarmInfo(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getAlarmInfo();
}

// icon activated callback
// ----------------------------------------------------------------------------
void SetIconActivatedCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setIconActivatedCallback(callback);
}

// Window
// ----------------------------------------------------------------------------
MainWindow_ *GetMainWindow() {
	return new MainWindow();
}

void MainWindowShow(MainWindow_ *w) {
	reinterpret_cast<MainWindow *>(w)->show();
	reinterpret_cast<MainWindow *>(w)->setFocus();
}

void MainWindowHidde(MainWindow_ *w) {
	reinterpret_cast<MainWindow *>(w)->hide();
}

void MainWindowUrl(MainWindow_ *w, char *url) {
	reinterpret_cast<MainWindow *>(w)->url(url);
}

void MainWindowFullScreen(MainWindow_ *w) {
	reinterpret_cast<MainWindow *>(w)->showFullScreen();
	reinterpret_cast<MainWindow *>(w)->setFocus();
}

void MainWindowNormal(MainWindow_ *w) {
	reinterpret_cast<MainWindow *>(w)->showNormal();
	reinterpret_cast<MainWindow *>(w)->setFocus();
}

void MainWindowThread(MainWindow_ *w, QThread_ *thread) {
	reinterpret_cast<MainWindow *>(w)->moveToThread(reinterpret_cast<QThread *>(thread));
}

void MainWindowDelete(MainWindow_ *w) {
	delete reinterpret_cast<MainWindow *>(w);
}

void MainWindowReload(MainWindow_ *w) {
	reinterpret_cast<MainWindow *>(w)->reload();
}

// lock screen
// ----------------------------------------------------------------------------
void SetLockScreen(SystemTray_ *t, int state) {
	reinterpret_cast<SystemTray *>(t)->setLockScreen(state, 0);
}

int GetLockScreen(SystemTray_ *t) {
	return reinterpret_cast<SystemTray *>(t)->getLockScreen();
}

void SetLockScreenCallback(SystemTray_ *t, void* callback) {
	reinterpret_cast<SystemTray *>(t)->setLockScreenCallback(callback);
}