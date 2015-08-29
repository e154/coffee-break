/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   mainwindow.h
 * Version:     v1.00
 * Created:     0:45:50 / 29 авг. 2015 г.
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

#ifndef MAINWINDOW_H_
#define MAINWINDOW_H_

#include <QWebView>
#include <QCloseEvent>

class MainWindow : public QWebView
{

//	Q_OBJECT

public:

	explicit MainWindow(QWidget *parent = 0);
	virtual ~MainWindow();

	void url(char *url);

private:
	void init(QWidget *parent);
	void closeEvent(QCloseEvent * event){ event->ignore(); }
};

#endif /* MAINWINDOW_H_ */
