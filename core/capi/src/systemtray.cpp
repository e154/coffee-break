/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   systemtray.cpp
 * Version:     v1.00
 * Created:     21:41:57 / 22 авг. 2015 г.
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


#include <stddef.h>
#include <iostream>

#include <QApplication>
#include <QVBoxLayout>
#include <QGroupBox>
#include <QCheckBox>
#include <QDebug>
#include <QAction>
#include <QActionGroup>
#include <QRadioButton>
#include <QObject>
#include <QWidget>
#include <QDialog>
#include <QMenu>
#include <QDesktopWidget>

#include "systemtray.h"
#include "capi.h"

//todo remove
#include <typeinfo>

SystemTray *SystemTray::mSystemTray = NULL;

SystemTray::SystemTray():
		mMainMenu(0),
		mAlarmMenu(0),
		mDefaultTimerMenu(0),

		mainLayout(0),
		exitAction(0),
		runAtStartUpAction(0),
		dTimerAction(0),
//		helpAction(0),
		alarmInfo(0),

		time4hAction(0),
		time3hAction(0),
		time2hAction(0),
		time1hAction(0),
		time45mAction(0),
		time30mAction(0),
		time25mAction(0),
		time20mAction(0),
		time15mAction(0),
		time10mAction(0),
		time5mAction(0),

		dTime4hAction(0),
		dTime3hAction(0),
		dTime2hAction(0),
		dTime1hAction(0),
		dTime45mAction(0),
		dTime30mAction(0),
		dTime25mAction(0),
		dTime20mAction(0),
		dTime15mAction(0),
		dTime10mAction(0),
		dTime5mAction(0),
		mTrayIcon(0),

		alarmAction1(0),
		alarmAction2(0),
		alarmAction3(0),
		mCurrentTimeLimit(0)
{
	init();
}

SystemTray* SystemTray::getSingletonPtr() {
    if( !mSystemTray ) {
    	mSystemTray = new SystemTray();
    }

    return mSystemTray;
}

//-------------------------------------------------------------------------------------
void SystemTray::init() {
	createActions();
	createTrayIcon();
}

void SystemTray::createTrayIcon() {

	mTrayIcon = new QSystemTrayIcon(this);

	connect(mTrayIcon, SIGNAL(activated(QSystemTrayIcon::ActivationReason)),
	            this, SLOT(iconActivated(QSystemTrayIcon::ActivationReason)));

	mMainMenu = new QMenu(this);

	connect(mMainMenu, SIGNAL(aboutToShow()), this, SLOT(trayAboutToShow()));
	trayAboutToShow();
	mTrayIcon->setContextMenu(mMainMenu);
}

void SystemTray::iconActivated(QSystemTrayIcon::ActivationReason reason)
{
    if(mIconActivatedCallback)
		go_callback_int(mIconActivatedCallback, reason);
}

void SystemTray::trayAboutToShow() {

	mMainMenu->clear();
	mainLayout = new QGridLayout;

	mMainMenu->addAction(exitAction);

	mMainMenu->addSeparator();
	mMainMenu->addAction(runAtStartUpAction);
	mDefaultTimerMenu = mMainMenu->addMenu(tr("&Default timer"));
	mAlarmMenu = mMainMenu->addMenu(tr("&Alarm sound"));
//	mMainMenu->addAction(helpAction);

	mMainMenu->addSeparator();
	mMainMenu->addAction(alarmInfo);
	mMainMenu->addSeparator();

	mMainMenu->addAction(time4hAction);
	mMainMenu->addAction(time3hAction);
	mMainMenu->addAction(time2hAction);
	mMainMenu->addAction(time1hAction);
	mMainMenu->addAction(time45mAction);
	mMainMenu->addAction(time30mAction);
	mMainMenu->addAction(time25mAction);
	mMainMenu->addAction(time20mAction);
	mMainMenu->addAction(time15mAction);
	mMainMenu->addAction(time10mAction);
	mMainMenu->addAction(time5mAction);

	// default time menu
	mDefaultTimerMenu->addAction(dTime4hAction);
	mDefaultTimerMenu->addAction(dTime3hAction);
	mDefaultTimerMenu->addAction(dTime2hAction);
	mDefaultTimerMenu->addAction(dTime45mAction);
	mDefaultTimerMenu->addAction(dTime30mAction);
	mDefaultTimerMenu->addAction(dTime25mAction);
	mDefaultTimerMenu->addAction(dTime20mAction);
	mDefaultTimerMenu->addAction(dTime15mAction);
	mDefaultTimerMenu->addAction(dTime10mAction);
	mDefaultTimerMenu->addAction(dTime5mAction);

	// alarm menu
	mAlarmMenu->addAction(alarmAction1);
//	mAlarmMenu->addAction(alarmAction2);
	mAlarmMenu->addAction(alarmAction3);

}

void SystemTray::createActions() {

	runAtStartUpAction = new QAction(tr("&Run at startup"), this);
	runAtStartUpAction->setCheckable(true);

//	helpAction = new QAction(tr("&Help"), this);

	alarmInfo = new QAction(tr("&Alarm is off"), this);
	alarmInfo->setDisabled(true);

	time4hAction = new QAction(tr("&4 hours"), this);
	time3hAction = new QAction(tr("&3 hours"), this);
	time2hAction = new QAction(tr("&2 hours"), this);
	time1hAction = new QAction(tr("&1 hour"), this);
	time45mAction = new QAction(tr("&45 minutes"), this);
	time30mAction = new QAction(tr("&30 minutes"), this);
	time25mAction = new QAction(tr("&25 minutes"), this);
	time20mAction = new QAction(tr("&20 minutes"), this);
	time15mAction = new QAction(tr("&15 minutes"), this);
	time10mAction = new QAction(tr("&10 minutes"), this);
	time5mAction = new QAction(tr("&5  minutes"), this);

	// time insert
	mTimeStates.insert( std::pair<int, QAction*>(4 * HOUR,time4hAction) );
	mTimeStates.insert( std::pair<int, QAction*>(3 * HOUR,time3hAction) );
	mTimeStates.insert( std::pair<int, QAction*>(2 * HOUR,time2hAction) );
	mTimeStates.insert( std::pair<int, QAction*>(1 * HOUR,time1hAction) );
	mTimeStates.insert( std::pair<int, QAction*>(45 * MINUTE,time45mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(30 * MINUTE,time30mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(25 * MINUTE,time25mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(20 * MINUTE,time20mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(15 * MINUTE,time15mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(10 * MINUTE,time10mAction) );
	mTimeStates.insert( std::pair<int, QAction*>(5 * MINUTE,time5mAction) );

	// default time menu
	dTime4hAction = new QAction(tr("&4 hours"), this);
	dTime3hAction = new QAction(tr("&3 hours"), this);
	dTime2hAction = new QAction(tr("&2 hours"), this);
	dTime1hAction = new QAction(tr("&1 hour"), this);
	dTime45mAction = new QAction(tr("&45 minutes"), this);
	dTime30mAction = new QAction(tr("&30 minutes"), this);
	dTime25mAction = new QAction(tr("&25 minutes"), this);
	dTime20mAction = new QAction(tr("&20 minutes"), this);
	dTime15mAction = new QAction(tr("&15 minutes"), this);
	dTime10mAction = new QAction(tr("&10 minutes"), this);
	dTime5mAction = new QAction(tr("&5  minutes"), this);

	mDtimeStates.insert( std::pair<int, QAction*>(4 * HOUR,dTime4hAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(3 * HOUR,dTime3hAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(2 * HOUR,dTime2hAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(1 * HOUR,dTime1hAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(45 * MINUTE,dTime45mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(30 * MINUTE,dTime30mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(25 * MINUTE,dTime25mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(20 * MINUTE,dTime20mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(15 * MINUTE,dTime15mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(10 * MINUTE,dTime10mAction) );
	mDtimeStates.insert( std::pair<int, QAction*>(5 * MINUTE,dTime5mAction) );

	// alarm menu
	alarmAction1 = new QAction(tr("&Alarm clock"), this);
//	alarmAction2 = new QAction(tr("A&nnoying alarm clock"), this);
	alarmAction3 = new QAction(tr("Di&sabled"), this);

	mAlarmStates.insert( std::pair<int, QAction*>(1,alarmAction1) );
//	mAlarmStates.insert( std::pair<int, QAction*>(2,alarmAction2) );
	mAlarmStates.insert( std::pair<int, QAction*>(3,alarmAction3) );

	connect(alarmAction1, SIGNAL(triggered()), this, SLOT(setAlarm1()));
//	connect(alarmAction2, SIGNAL(triggered()), this, SLOT(setAlarm2()));
	connect(alarmAction3, SIGNAL(triggered()), this, SLOT(setAlarm3()));

	exitAction = new QAction(tr("&Exit"), this);

	// quit
	connect(exitAction, SIGNAL(triggered()), qApp, SLOT(quit()));

	// run At StartUp
	connect(runAtStartUpAction, SIGNAL(triggered()), this, SLOT(setRunAtStartup1()));

	// show help
//	connect(helpAction, SIGNAL(triggered()), this, SLOT(showHelp()));

	// set default time
	connect(dTime4hAction, SIGNAL(triggered()), this, SLOT(set4hdTime()));
	connect(dTime3hAction, SIGNAL(triggered()), this, SLOT(set3hdTime()));
	connect(dTime2hAction, SIGNAL(triggered()), this, SLOT(set2hdTime()));
	connect(dTime1hAction, SIGNAL(triggered()), this, SLOT(set1hdTime()));
	connect(dTime45mAction, SIGNAL(triggered()), this, SLOT(set45mdTime()));
	connect(dTime30mAction, SIGNAL(triggered()), this, SLOT(set30mdTime()));
	connect(dTime25mAction, SIGNAL(triggered()), this, SLOT(set25mdTime()));
	connect(dTime20mAction, SIGNAL(triggered()), this, SLOT(set20mdTime()));
	connect(dTime15mAction, SIGNAL(triggered()), this, SLOT(set15mdTime()));
	connect(dTime10mAction, SIGNAL(triggered()), this, SLOT(set10mdTime()));
	connect(dTime5mAction, SIGNAL(triggered()), this, SLOT(set5mdTime()));

	// set time
	connect(time4hAction, SIGNAL(triggered()), this, SLOT(set4hTime()));
	connect(time3hAction, SIGNAL(triggered()), this, SLOT(set3hTime()));
	connect(time2hAction, SIGNAL(triggered()), this, SLOT(set2hTime()));
	connect(time1hAction, SIGNAL(triggered()), this, SLOT(set1hTime()));
	connect(time45mAction, SIGNAL(triggered()), this, SLOT(set45mTime()));
	connect(time30mAction, SIGNAL(triggered()), this, SLOT(set30mTime()));
	connect(time25mAction, SIGNAL(triggered()), this, SLOT(set25mTime()));
	connect(time20mAction, SIGNAL(triggered()), this, SLOT(set20mTime()));
	connect(time15mAction, SIGNAL(triggered()), this, SLOT(set15mTime()));
	connect(time10mAction, SIGNAL(triggered()), this, SLOT(set10mTime()));
	connect(time5mAction, SIGNAL(triggered()), this, SLOT(set5mTime()));
}

void SystemTray::setDTimer(int inTime, QAction *action) {

	int time = 45 * MINUTE;
	if ( !mDtimeStates.empty() ) {
		actionStates::iterator it = mDtimeStates.begin();
		while ( it != mDtimeStates.end()){

			if (
					(inTime == 0 && action != 0x0 && it->second != action) ||
					(inTime != 0 && action == 0x0 && it->first != inTime)
					)
			{
				it->second->setDisabled(false);
				it->second->setChecked(false);
				it->second->setCheckable(false);

			} else {

				it->second->setDisabled(true);
				it->second->setCheckable(true);
				it->second->setChecked(true);
				time = it->first;
			}
			++it;
		}
	}

	mCurrentDefaultTime = (inTime == 0) ? time : inTime;

	// call back to go side
	if(mDtimeCallback)
		go_callback_int(mDtimeCallback, time);
}

void SystemTray::setAlarm(int inState, QAction *action) {

	int state = 1;
	if ( !mAlarmStates.empty() ) {
		actionStates::iterator it = mAlarmStates.begin();
		while ( it != mAlarmStates.end()){

			if (
					(inState == 0 && action != 0x0 && it->second != action) ||
					(inState != 0 && action == 0x0 && it->first != inState)
					)
			{
				it->second->setDisabled(false);
				it->second->setChecked(false);
				it->second->setCheckable(false);

			} else {

				it->second->setDisabled(true);
				it->second->setCheckable(true);
				it->second->setChecked(true);
				state = it->first;
			}
			++it;
		}
	}

	mCurrentAlarm = (inState == 0) ? state : inState;

	// call back to go side
	if(mAlarmCallback)
		go_callback_int(mAlarmCallback, state);
}

void SystemTray::setTimer(int inTime, QAction *action) {

	int time = 45 * MINUTE;
	if ( !mTimeStates.empty() ) {
		actionStates::iterator it = mTimeStates.begin();
		while ( it != mTimeStates.end()){

			if (
					(inTime == 0 && action != 0x0 && it->second != action) ||
					(inTime != 0 && action == 0x0 && it->first != inTime)
					)
			{
				it->second->setDisabled(false);
				it->second->setChecked(false);
				it->second->setCheckable(false);

			} else {

				it->second->setDisabled(true);
				it->second->setCheckable(true);
				it->second->setChecked(true);
				time = it->first;
			}
			++it;
		}
	}

	mCurrentTimeLimit = (inTime == 0) ? time : inTime;

	// call back to go side
	if(mTimeCallback)
		go_callback_int(mTimeCallback, time);
}

void SystemTray::setRunAtStartup(int state) {

	if(state) {
		runAtStartUpAction->setChecked(true);
		mCurrentRunAtStartup = 1;
	} else {
		runAtStartUpAction->setChecked(false);
		mCurrentRunAtStartup = 0;
	}

	// call back to go side
	if(mRunAtStartupCallback)
		go_callback_int(mRunAtStartupCallback, mCurrentRunAtStartup);
}
