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

#include "systemtray.h"
#include "capi.h"

SystemTray *SystemTray::mSystemTray = NULL;

SystemTray::SystemTray():
		mMainMenu(0),
		mAlarmMenu(0),
		mTimerMenu(0),

		mainLayout(0),
		exitAction(0),
		runAtStartUpAction(0),
		dTimerAction(0),
		helpAction(0),
		alarmInfo(0),

		timeLayout(0),
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

		dTimeLayout(0),
		dTime4hRadio(0),
		dTime3hRadio(0),
		dTime2hRadio(0),
		dTime1hRadio(0),
		dTime45mRadio(0),
		dTime30mRadio(0),
		dTime25mRadio(0),
		dTime20mRadio(0),
		dTime15mRadio(0),
		dTime10mRadio(0),
		dTime5mRadio(0),
		mTrayIcon(0),
		alarmLayout(0),
		alarmRadio1(0),
		alarmRadio2(0),
		alarmRadio3(0),
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
    switch (reason) {
    case QSystemTrayIcon::Trigger:
    case QSystemTrayIcon::DoubleClick:
    	showMessage("title", "message");
        qDebug() << "dblclick";
        break;
    default:
        ;
    }
}

void SystemTray::trayAboutToShow() {

	mMainMenu->clear();
	mainLayout = new QGridLayout;

	mMainMenu->addAction(exitAction);

	mMainMenu->addSeparator();
	mMainMenu->addAction(runAtStartUpAction);
	mTimerMenu = mMainMenu->addMenu(tr("&Default timer"));
	mAlarmMenu = mMainMenu->addMenu(tr("&Alarm sound"));
	mMainMenu->addAction(helpAction);

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
	dTimeLayout = new QGridLayout;
	dTimeLayout->addWidget(dTime4hRadio);
	dTimeLayout->addWidget(dTime3hRadio);
	dTimeLayout->addWidget(dTime2hRadio);
	dTimeLayout->addWidget(dTime45mRadio);
	dTimeLayout->addWidget(dTime30mRadio);
	dTimeLayout->addWidget(dTime25mRadio);
	dTimeLayout->addWidget(dTime20mRadio);
	dTimeLayout->addWidget(dTime15mRadio);
	dTimeLayout->addWidget(dTime10mRadio);
	dTimeLayout->addWidget(dTime5mRadio);
	mTimerMenu->setLayout(dTimeLayout);

	// alarm menu
	alarmLayout = new QGridLayout;
	alarmLayout->addWidget(alarmRadio1);
	alarmLayout->addWidget(alarmRadio2);
	alarmLayout->addWidget(alarmRadio3);
	mAlarmMenu->setLayout(alarmLayout);
}

void SystemTray::createActions() {

	runAtStartUpAction = new QAction(tr("&Run at startup"), this);
	runAtStartUpAction->setCheckable(true);

	helpAction = new QAction(tr("&Help"), this);

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

	//default time menu
	dTime4hRadio = new QRadioButton(tr("&4 hours"), this);
	dTime3hRadio = new QRadioButton(tr("&3 hours"), this);
	dTime2hRadio = new QRadioButton(tr("&2 hours"), this);
	dTime1hRadio = new QRadioButton(tr("&1 hour"), this);
	dTime45mRadio = new QRadioButton(tr("&45 minutes"), this);
	dTime30mRadio = new QRadioButton(tr("&30 minutes"), this);
	dTime25mRadio = new QRadioButton(tr("&25 minutes"), this);
	dTime20mRadio = new QRadioButton(tr("&20 minutes"), this);
	dTime15mRadio = new QRadioButton(tr("&15 minutes"), this);
	dTime10mRadio = new QRadioButton(tr("&10 minutes"), this);
	dTime5mRadio = new QRadioButton(tr("&5  minutes"), this);

	alarmRadio1 = new QRadioButton(tr("&Alarm clock"), this);
	alarmRadio2 = new QRadioButton(tr("A&nnoying alarm clock"), this);
	alarmRadio3 = new QRadioButton(tr("Di&sabled"), this);

	exitAction = new QAction(tr("&Exit"), this);

	// quit
	connect(exitAction, SIGNAL(triggered()), qApp, SLOT(quit()));

	// run At StartUp
	connect(runAtStartUpAction, SIGNAL(triggered()), this, SLOT(setRunAtStartUp()));

	// show help
	connect(helpAction, SIGNAL(triggered()), this, SLOT(showHelp()));

	// set alarm
	connect(alarmRadio1, &QAbstractButton::toggled, this, [=](int state){ setAlarm(state, 1); });
	connect(alarmRadio2, &QAbstractButton::toggled, this, [=](int state){ setAlarm(state, 2); });
	connect(alarmRadio3, &QAbstractButton::toggled, this, [=](int state){ setAlarm(state, 3); });

	// set default time
	connect(dTime4hRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 4 * HOUR); });
	connect(dTime3hRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 3 * HOUR); });
	connect(dTime2hRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 2 * HOUR); });
	connect(dTime1hRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 1 * HOUR); });
	connect(dTime45mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 45 * MINUTE); });
	connect(dTime30mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 30 * MINUTE); });
	connect(dTime25mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 25 * MINUTE); });
	connect(dTime20mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 20 * MINUTE); });
	connect(dTime15mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 15 * MINUTE); });
	connect(dTime10mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 10 * MINUTE); });
	connect(dTime5mRadio, &QAbstractButton::toggled, this, [=](int state){ setDefaultTimer(state, 5 * MINUTE); });

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

void SystemTray::setDefaultTimer(int state, int value) {

	if(state == 0) return;

	qDebug() << "set default time: " << value;
}

void SystemTray::setAlarm(int state, int value) {

	if(state == 0) return;

	qDebug() << "set alarm: " << value;
}

void SystemTray::showHelp() {

	qDebug() << "show help";
}

void SystemTray::setTimer(QAction *action) {

	int time;
	if ( !mTimeStates.empty() ) {
		actionStates::iterator it = mTimeStates.begin();
		while ( it != mTimeStates.end()){
			if ( it->second != action) {
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

	mCurrentTimeLimit = time;

	// call back to go program
	if(mTimeCallback)
		go_callback_int(mTimeCallback, time);
}

void SystemTray::setTimer(int time) {

	if ( !mTimeStates.empty() ) {
		actionStates::iterator it = mTimeStates.begin();
		while ( it != mTimeStates.end()){
			if ( it->first != time) {
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

	mCurrentTimeLimit = time;
}

void SystemTray::setRunAtStartUp() {

	qDebug() << "Run at startup: " << runAtStartUpAction->isChecked();
}
