/****************************************************************************
** Meta object code from reading C++ file 'systemtray.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.3.2)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "systemtray.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'systemtray.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.3.2. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
struct qt_meta_stringdata_SystemTray_t {
    QByteArrayData data[23];
    char stringdata[261];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_SystemTray_t, stringdata) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_SystemTray_t qt_meta_stringdata_SystemTray = {
    {
QT_MOC_LITERAL(0, 0, 10),
QT_MOC_LITERAL(1, 11, 15),
QT_MOC_LITERAL(2, 27, 0),
QT_MOC_LITERAL(3, 28, 13),
QT_MOC_LITERAL(4, 42, 33),
QT_MOC_LITERAL(5, 76, 6),
QT_MOC_LITERAL(6, 83, 15),
QT_MOC_LITERAL(7, 99, 5),
QT_MOC_LITERAL(8, 105, 5),
QT_MOC_LITERAL(9, 111, 8),
QT_MOC_LITERAL(10, 120, 8),
QT_MOC_LITERAL(11, 129, 15),
QT_MOC_LITERAL(12, 145, 9),
QT_MOC_LITERAL(13, 155, 9),
QT_MOC_LITERAL(14, 165, 9),
QT_MOC_LITERAL(15, 175, 9),
QT_MOC_LITERAL(16, 185, 10),
QT_MOC_LITERAL(17, 196, 10),
QT_MOC_LITERAL(18, 207, 10),
QT_MOC_LITERAL(19, 218, 10),
QT_MOC_LITERAL(20, 229, 10),
QT_MOC_LITERAL(21, 240, 10),
QT_MOC_LITERAL(22, 251, 9)
    },
    "SystemTray\0trayAboutToShow\0\0iconActivated\0"
    "QSystemTrayIcon::ActivationReason\0"
    "reason\0setDefaultTimer\0state\0value\0"
    "setAlarm\0showHelp\0setRunAtStartUp\0"
    "set4hTime\0set3hTime\0set2hTime\0set1hTime\0"
    "set45mTime\0set30mTime\0set25mTime\0"
    "set20mTime\0set15mTime\0set10mTime\0"
    "set5mTime"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_SystemTray[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      17,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    0,   99,    2, 0x08 /* Private */,
       3,    1,  100,    2, 0x08 /* Private */,
       6,    2,  103,    2, 0x08 /* Private */,
       9,    2,  108,    2, 0x08 /* Private */,
      10,    0,  113,    2, 0x08 /* Private */,
      11,    0,  114,    2, 0x08 /* Private */,
      12,    0,  115,    2, 0x08 /* Private */,
      13,    0,  116,    2, 0x08 /* Private */,
      14,    0,  117,    2, 0x08 /* Private */,
      15,    0,  118,    2, 0x08 /* Private */,
      16,    0,  119,    2, 0x08 /* Private */,
      17,    0,  120,    2, 0x08 /* Private */,
      18,    0,  121,    2, 0x08 /* Private */,
      19,    0,  122,    2, 0x08 /* Private */,
      20,    0,  123,    2, 0x08 /* Private */,
      21,    0,  124,    2, 0x08 /* Private */,
      22,    0,  125,    2, 0x08 /* Private */,

 // slots: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,
    QMetaType::Void, QMetaType::Int, QMetaType::Int,    7,    8,
    QMetaType::Void, QMetaType::Int, QMetaType::Int,    7,    8,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,

       0        // eod
};

void SystemTray::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        SystemTray *_t = static_cast<SystemTray *>(_o);
        switch (_id) {
        case 0: _t->trayAboutToShow(); break;
        case 1: _t->iconActivated((*reinterpret_cast< QSystemTrayIcon::ActivationReason(*)>(_a[1]))); break;
        case 2: _t->setDefaultTimer((*reinterpret_cast< int(*)>(_a[1])),(*reinterpret_cast< int(*)>(_a[2]))); break;
        case 3: _t->setAlarm((*reinterpret_cast< int(*)>(_a[1])),(*reinterpret_cast< int(*)>(_a[2]))); break;
        case 4: _t->showHelp(); break;
        case 5: _t->setRunAtStartUp(); break;
        case 6: _t->set4hTime(); break;
        case 7: _t->set3hTime(); break;
        case 8: _t->set2hTime(); break;
        case 9: _t->set1hTime(); break;
        case 10: _t->set45mTime(); break;
        case 11: _t->set30mTime(); break;
        case 12: _t->set25mTime(); break;
        case 13: _t->set20mTime(); break;
        case 14: _t->set15mTime(); break;
        case 15: _t->set10mTime(); break;
        case 16: _t->set5mTime(); break;
        default: ;
        }
    }
}

const QMetaObject SystemTray::staticMetaObject = {
    { &QDialog::staticMetaObject, qt_meta_stringdata_SystemTray.data,
      qt_meta_data_SystemTray,  qt_static_metacall, 0, 0}
};


const QMetaObject *SystemTray::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *SystemTray::qt_metacast(const char *_clname)
{
    if (!_clname) return 0;
    if (!strcmp(_clname, qt_meta_stringdata_SystemTray.stringdata))
        return static_cast<void*>(const_cast< SystemTray*>(this));
    return QDialog::qt_metacast(_clname);
}

int SystemTray::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QDialog::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 17)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 17;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 17)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 17;
    }
    return _id;
}
QT_END_MOC_NAMESPACE
