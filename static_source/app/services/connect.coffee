'use strict'

angular
  .module('appServices')
  .service 'Connect', ()->
    #--------------------------------------------------------
    # socket
    #--------------------------------------------------------
    if !WebSocket in window
      return

    port = document.location.port
    protocol = if document.location.protocol == "https:" then "wss:" else "ws:"
    socket = new WebSocket(protocol+"//"+document.domain+":"+port+"/ws")

    socket.onclose = ()=>
      console.log("socket closed")

    socket.onopen = ()=>
      console.log("socket opened")

    socket.onmessage = (e)=>
      console.log("socket message")
