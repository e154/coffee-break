'use strict'

angular
.module('app')
.directive 'typetype', ['$window', '$timeout'
  ($window, $timeout)=>
    restrict: 'AE'
    scope: {
      text: "="
      options: "="
      ngModel: "="
    }
    link: ($scope, $element, $attrs)=>

      el = $($element)

      options = angular.extend $scope.options,
        keypress: ()->
          update()
        callback: ()->
          update()

      update = ()->
        $timeout ()->
          $scope.$apply(
            $scope.ngModel = {val: el.val()}
          )
        , 0

      el.typetype($scope.text, options)
  ]