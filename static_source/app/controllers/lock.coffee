'use strict'

angular
  .module('appControllers')
  .controller 'lockCtrl', ['$scope', 'ngSocket'
  ($scope, ngSocket) ->
    vm = this
    vm.uptime_total = 0
    vm.uptime_idle = 0
    vm.lock = 0
    vm.lock_const = 0
    vm.total_work = 0
    vm.total_idle = 0
    vm.unlock_wait = 0

    port = document.location.port
    protocol = if document.location.protocol == "https:" then "wss:" else "ws:"
    ws = ngSocket(protocol+"//"+document.domain+":"+port+"/ws")

    toSeconds = (val)->
      if val?
        val / 1000000000

    toHumanTime = (sec_num)->

      time = ''
      if sec_num < 0
        sec_num *= -1
        time += '-'

      hours   = Math.floor(sec_num / 3600)
      minutes = Math.floor((sec_num - (hours * 3600)) / 60)
      seconds = sec_num - (hours * 3600) - (minutes * 60)

      if hours < 10
        hours = '0' + hours

      if minutes < 10
        minutes = '0' + minutes

      if seconds < 10
        seconds = '0' + seconds

      if hours != '00'
        time += hours+ 'h'

      time += minutes + 'm' + seconds + 's'

      time

    ws.onMessage (message)=>

      data = angular.fromJson(message.data)

      if data.uptime_total?
        vm.uptime_total = data.uptime_total

      if data.uptime_idle?
        vm.uptime_idle = data.uptime_idle

      if data.timeinfo?
        if data.timeinfo['lock']?
          vm.lock = toHumanTime(toSeconds(data.timeinfo['lock']))

        if data.timeinfo.lock_const?
          vm.lock_const = toHumanTime(toSeconds(data.timeinfo.lock_const))

        if data.timeinfo.total_work?
          vm.total_work = toHumanTime(toSeconds(data.timeinfo.total_work))

        if data.timeinfo.total_idle?
          vm.total_idle = toHumanTime(toSeconds(data.timeinfo.total_idle))

        if data.timeinfo.lock_const && data.timeinfo.lock?
          vm.unlock_wait = toHumanTime(toSeconds(data.timeinfo.lock_const - data.timeinfo.lock))
  ]