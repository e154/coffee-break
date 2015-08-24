import QtQuick 2.0

Rectangle{
    width: 360
    height: 360
    color: "grey"

    Button{
        id: loadButton
        label: "Load"
        onButtonClick: {
            console.log("Hello!")
        }
    }
}
