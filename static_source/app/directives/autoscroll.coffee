'use strict'

angular
.module('app')
.directive 'autoscroll', ['$window', '$timeout'
  ($window, $timeout)=>
    restrict: 'AE'
    link: ($scope, $element, $attrs)=>

      el = $($element)

      $window.setInterval ()->
        el.scrollTop(el[0].scrollHeight)
      , 1000
  ]