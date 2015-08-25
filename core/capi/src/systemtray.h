/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   systemtray.h
 * Version:     v1.00
 * Created:     21:41:56 / 22 авг. 2015 г.
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

#ifndef SYSTEMTRAY_H_
#define SYSTEMTRAY_H_

#include <QSystemTrayIcon>
#include <QGridLayout>
#include <QDialog>
#include <QAction>
#include <iostream>
#include <map>

const int SECOND = 1;
const int MINUTE = 60 * SECOND;
const int HOUR = 60 * MINUTE;

class QRadioButton;

typedef std::map<int, QAction*> actionStates;

class SystemTray : public QDialog
{
	Q_OBJECT

private:
	QMenu *mMainMenu;
	QMenu *mAlarmMenu;
	QMenu *mDefaultTimerMenu;

	QGridLayout *mainLayout;
	QAction *exitAction;
	QAction *runAtStartUpAction;
	QAction *dTimerAction;
	QAction *AlarmSoundAction;
	QAction *helpAction;
	QAction *alarmInfo;

	QAction *time4hAction;
	QAction *time3hAction;
	QAction *time2hAction;
	QAction *time1hAction;
	QAction *time45mAction;
	QAction *time30mAction;
	QAction *time25mAction;
	QAction *time20mAction;
	QAction *time15mAction;
	QAction *time10mAction;
	QAction *time5mAction;

	QAction *dTime4hAction;
	QAction *dTime3hAction;
	QAction *dTime2hAction;
	QAction *dTime1hAction;
	QAction *dTime45mAction;
	QAction *dTime30mAction;
	QAction *dTime25mAction;
	QAction *dTime20mAction;
	QAction *dTime15mAction;
	QAction *dTime10mAction;
	QAction *dTime5mAction;

	QSystemTrayIcon *mTrayIcon;

	QAction *alarmAction1;
	QAction *alarmAction2;
	QAction *alarmAction3;

public:
	SystemTray();

	void init();

	inline QSystemTrayIcon *GetTrayPtr() { return mTrayIcon; }
	inline QMenu *GetMenuPtr() { return mMainMenu; }

	static SystemTray* getSingletonPtr();
	void setTrayIcon(char *icon) { mTrayIcon->setIcon(QIcon(icon)); }
	void setTrayToolTip(char *tooltip) { mTrayIcon->setToolTip(QString(tooltip)); }
	void setTrayVisible(bool trigger) { mTrayIcon->setVisible(trigger); }
	void showMessage(const QString &title, const QString &msg,
	                     QSystemTrayIcon::MessageIcon icon = QSystemTrayIcon::Information, int msecs = 10000) {
		mTrayIcon->showMessage(title, msg, icon, msecs);
	}

	// time
	void setTimer(int time = 0, QAction *action = 0);
	int getTime() { return mCurrentTimeLimit; }
	void setTimeCallback(void* callback) {
		mTimeCallback = callback;
	}

	// default time
	void setDTimer(int time = 0, QAction *action = 0);
	int getDtime() { return mCurrentDefaultTime; }
	void setDtimeCallback(void* callback) {
		mDtimeCallback = callback;
	}

	// alarm state
	void setAlarm(int inState, QAction *action);
	int getAlarm() { return mCurrentAlarm; }
	void setAlarmCallback(void* callback) {
		mDtimeCallback = callback;
	}

	// run at startup
	void setRunAtStartup(int state = 0);
	int getRunAtStartup() { return mCurrentRunAtStartup; }
	void setRunAtStartupCallback(void* callback) {
		mRunAtStartupCallback = callback;
	}

	// alarm info
	void setAlarmInfo(char *text) { alarmInfo->setText(QString(text)); }
	char *getAlarmInfo() { return (char*)alarmInfo->text().data(); }

	// icon activated callback
	void setIconActivatedCallback(void* callback) {
		mIconActivatedCallback = callback;
	}

private slots:
	void trayAboutToShow();
	void iconActivated(QSystemTrayIcon::ActivationReason reason);
	void showHelp() {};

	// time
	inline void set4hTime() { setTimer(0, time4hAction); }
	inline void set3hTime() { setTimer(0, time3hAction); }
	inline void set2hTime() { setTimer(0, time2hAction); }
	inline void set1hTime() { setTimer(0, time1hAction); }
	inline void set45mTime() { setTimer(0, time45mAction); }
	inline void set30mTime() { setTimer(0, time30mAction); }
	inline void set25mTime() { setTimer(0, time25mAction); }
	inline void set20mTime() { setTimer(0, time20mAction); }
	inline void set15mTime() { setTimer(0, time15mAction); }
	inline void set10mTime() { setTimer(0, time10mAction); }
	inline void set5mTime() { setTimer(0, time5mAction); }

	// default time
	inline void set4hdTime() { setDTimer(0, dTime4hAction); }
	inline void set3hdTime() { setDTimer(0, dTime3hAction); }
	inline void set2hdTime() { setDTimer(0, dTime2hAction); }
	inline void set1hdTime() { setDTimer(0, dTime1hAction); }
	inline void set45mdTime() { setDTimer(0, dTime45mAction); }
	inline void set30mdTime() { setDTimer(0, dTime30mAction); }
	inline void set25mdTime() { setDTimer(0, dTime25mAction); }
	inline void set20mdTime() { setDTimer(0, dTime20mAction); }
	inline void set15mdTime() { setDTimer(0, dTime15mAction); }
	inline void set10mdTime() { setDTimer(0, dTime10mAction); }
	inline void set5mdTime() { setDTimer(0, dTime5mAction); }

	// alarm
	inline void setAlarm1() { setAlarm(0, alarmAction1); }
	inline void setAlarm2() { setAlarm(0, alarmAction2); }
	inline void setAlarm3() { setAlarm(0, alarmAction3); }

	inline void setRunAtStartup1() {

		int state = (runAtStartUpAction->isChecked())?1:0;
		setRunAtStartup(state);
	};

private:
	void createTrayIcon();
	void createActions();

	static SystemTray *mSystemTray;

protected:
	void* mTimeCallback;
	void* mDtimeCallback;
	void* mAlarmCallback;
	void* mRunAtStartupCallback;
	void* mIconActivatedCallback;
	actionStates mTimeStates;
	actionStates mDtimeStates;
	actionStates mAlarmStates;
	int mCurrentTimeLimit;
	int mCurrentDefaultTime;
	int mCurrentAlarm;
	int mCurrentRunAtStartup;

};

#endif /* SYSTEMTRAY_H_ */
