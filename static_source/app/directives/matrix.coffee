#'use strict'

angular
  .module('app')
  .directive 'matrix', ['$window', '$document'
  ($window, $document)=>
    restrict: 'A'
    link: ($scope, $element, $attrs)=>

      canvas = {}
      lienzo = {}
      tam_letra = 16
      columna = new Array()

      DOMContentLoaded = ()=>
        canvas = document.getElementById('contenedor')
        canvas.height = $window.innerHeight - 5
        canvas.width = $window.innerWidth

        lienzo = canvas.getContext('2d')
        total_columnas =  Math.round(canvas.width / tam_letra)

        for i in [0..total_columnas]
          columna[i] = 1

        setInterval(dibujar, 85)
        $window.removeEventListener 'resize', recargar
        $window.addEventListener 'resize', recargar

      dibujar = ()=>
        lienzo.fillStyle = "rgba(0, 0, 0, 0.05)"
        lienzo.fillRect(0, 0, canvas.width, canvas.height)
        lienzo.fillStyle = "#00FF00"
        lienzo.font = tam_letra + "px Arial"

        for i in [0..columna.length]
          lienzo.fillText(texto_aleatorio(), (i*tam_letra), (columna[i]*tam_letra))
          if (columna[i] * tam_letra) > canvas.height && Math.random() > 0.975
            columna[i] = 0
          columna[i]++

      texto_aleatorio = ()=>
        String.fromCharCode(parseInt(Math.floor((Math.random() * 94) + 33)))

      recargar = ()=>
        canvas.height = $window.innerHeight - 5
        canvas.width = $window.innerWidth

        total_columnas =  Math.round(canvas.width / tam_letra)
        columna = new Array()

        for i in [0..total_columnas]
          columna[i] = 1

      DOMContentLoaded()
  ]