/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   capi.h
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

#ifndef CAPI_CAPI_H_
#define CAPI_CAPI_H_

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

extern void go_callback_int(void* foo, int p1);

enum MessageIcon { NoIcon, Information, Warning, Critical };

typedef void SystemTray_;
typedef void QApp_;
typedef void QThread_;
typedef void MainWindow_;

// Application
void NewGuiApplication();
void ApplicationExec();
void ApplicationExit();
void ApplicationFlushAll();
QApp_ *ApplicationPtr();
QThread_ *ApplicationThread();

// QSystemTrayIcon
SystemTray_ *GetSystemTray();
void SetTrayIcon(SystemTray_*, char *img);
void SetTrayToolTip(SystemTray_*, char *tooltip);
void SetTrayVisible(SystemTray_*, bool trigger);
void ShowMessage(SystemTray_*, char *title, char *msg, int icon);
void MoveToThread(SystemTray_*, QThread_ *thread);

// time
void SetTime(SystemTray_*, int time);
int GetTime(SystemTray_*);
void SetTimeCallback(SystemTray_*, void*);

// default time
void SetDTime(SystemTray_*, int time);
int GetDTime(SystemTray_*);
void SetDTimeCallback(SystemTray_*, void*);

// alert
void SetAlarm(SystemTray_*, int time);
int GetAlarm(SystemTray_*);
void SetAlarmCallback(SystemTray_*, void*);

// run at startup
void SetRunAtStartup(SystemTray_*, int time);
int GetRunAtStartup(SystemTray_*);
void SetRunAtStartupCallback(SystemTray_*, void*);

// alarm info
void SetAlarmInfo(SystemTray_*, char *time);
char *GetAlarmInfo(SystemTray_*);

void SetIconActivatedCallback(SystemTray_*, void*);

// MainWindow
MainWindow_ *GetMainWindow();
void MainWindowShow(MainWindow_ *w);
void MainWindowHidde(MainWindow_ *w);
void MainWindowUrl(MainWindow_ *w, char *url);

#ifdef __cplusplus
} // extern "C"
#endif

#endif /* CAPI_CAPI_H_ */
