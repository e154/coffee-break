/** Copyright (C), DeltaSync Studios, 2010-2014. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   mainwindow.cpp
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

#include <QWebView>
#include "mainwindow.h"

MainWindow::MainWindow(QWidget *parent) :
	QWebView(parent)
{
	init(parent);
}

MainWindow::~MainWindow() {

}

void MainWindow::init(QWidget *parent) {

	QPalette palette = this->palette();
    palette.setBrush(QPalette::Base, Qt::transparent);
    page()->setPalette(palette);
    setAttribute(Qt::WA_OpaquePaintEvent, false);
    setAttribute(Qt::WA_TranslucentBackground, true);
    setWindowFlags(Qt::FramelessWindowHint | Qt::Tool);

    setFocus();
}

void MainWindow::url(char *url) {
	load(QUrl(url));
}

