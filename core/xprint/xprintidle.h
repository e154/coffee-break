/** Copyright (C), DeltaSync Studios, 2010-2015. All rights reserved.
 * ------------------------------------------------------------------
 * File name:   xprintidle.h
 * Version:     v1.00
 * Created:     21:42:56 / 08 авг. 2015 г.
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

#ifndef XPRINTIDLE_H_
#define XPRINTIDLE_H_

#include <X11/Xlib.h>
#include <X11/extensions/dpms.h>
#include <X11/extensions/scrnsaver.h>

// errors
// 1 couldn't open display
// 2 screen saver extension not supported
// 3 couldn't query screen saver info

int getIdle( unsigned long *idle );
unsigned long workaroundCreepyXServer(Display *dpy, unsigned long _idleTime );


#endif /* XPRINTIDLE_H_ */
