'use strict'

angular
  .module('appControllers')
  .controller 'lockCtrl', ['$scope', 'ngSocket'
  ($scope, ngSocket) ->
    vm = this

    port = document.location.port
    protocol = if document.location.protocol == "https:" then "wss:" else "ws:"
    ws = ngSocket(protocol+"//"+document.domain+":"+port+"/ws")

    ws.onMessage (message)=>

      data = angular.fromJson(message.data)

      if data.result?
        console.log(data.result)


  ]