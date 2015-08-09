
'use strict'

angular.module('appFilters', [])
angular.module('appControllers', [])
angular.module('appServices', ['ngResource'])
app = angular
  .module('app', [
    'ngRoute'
    'appControllers'
    'appFilters'
    'appServices'
    'route-segment'
    'view-segment'
  ])

angular.module('app')
  .config ['$routeProvider', '$locationProvider', '$routeSegmentProvider',
  ($routeProvider, $locationProvider, $routeSegmentProvider) ->
    $routeSegmentProvider
      .when '/', 'base.dashboard'
      .when '/settings', 'base.settings'

      .segment 'base',
        templateUrl: '/templates/base.html'
        controller: 'baseCtrl as base'

      .within()
        .segment 'dashboard',
          default: true
          templateUrl: '/templates/dashboard.html'
          controller: 'dashboardCtrl as dashboard'

        .segment 'settings',
          templateUrl: '/templates/settings.html'
          controller: 'settingsCtrl as settings'

    $routeProvider.otherwise
      redirectTo: '/'
  ]

angular.module('app')
  .run ['$rootScope'
  ($rootScope) =>

    gui = require('nw.gui')
    win = gui.Window.get()
    tray

    win.on 'minimize', () ->
      this.hide()
      tray = new gui.Tray
        icon: '/images/icon.png'
      tray.on 'click', () ->
        win.show()
        this.remove()
        tray = null
  ]