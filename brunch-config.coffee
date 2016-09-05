exports.config =
  # See http://brunch.io/#documentation for docs.
  files:
    javascripts:
      joinTo: 'app.js'
    stylesheets:
      joinTo: 'app.css'
    templates:
      joinTo: 'index.html'
  npm:
    globals:
      THREE: 'three'
      dat: 'dat-gui'
  plugins:
    babel:
      pattern: /\.(es6|jsx|js)$/
      presets: [ 'es2015', 'es2016' , 'react']
      ignore: [
        /^vendor/
      ]
    sass:
      mode: 'native'
      allowCache: true
      options: ['--quiet']
      sourceMapEmbed: false
      includePaths: [
        'node_modules/bourbon/app/assets/stylesheets',
        'node_modules/normalize.css/normalize.css'
      ]
    jaded:
      staticPatterns: /^app(\/|\\)pug(\/|\\)(.+)\.pug$/
      jade:
        doctype: 'html'
        basedir: 'pug'
        filename: 'index'
        cache: true
        pretty: yes