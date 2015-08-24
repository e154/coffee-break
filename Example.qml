import QtQuick 2.0

Rectangle {
    id: button
    radius: 6
    border.width: 3
    border.color: "#ffffff"
    width: 150; height: 75
    property string label: ""
    color: "#eeeeee"
    signal buttonClick()
    onButtonClick: {
    }

    Text{
        id: buttonLabel
        anchors.centerIn: parent
        text: label
    }

    MouseArea {
        id: buttonMouseArea

        anchors.fill: parent
        onClicked: buttonClick()
    }
}