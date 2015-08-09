var source = "static_source";
var pub = "public";
var tmp = "tmp";

module.exports = {
    build_lib_js: {
        filename: 'lib.min.js',
        paths: {
            bowerDirectory: source + '/bower_components',
            bowerrc: '.bowerrc',
            bowerJson: 'bower.json'
        },
        dest: pub + '/js'
    },
    build_coffee_js: {
        filename: 'app.min.js',
        source: [
            "app/assets/javascripts/modules.coffee",
            "app/assets/javascripts/language.coffee",
            "app/assets/javascripts/routes.coffee",
            "app/assets/javascripts/setup.coffee",
            "app/assets/javascripts/fixes.coffee",
            "app/assets/javascripts/auth.coffee",
            "app/assets/javascripts/stream.coffee",
            "app/assets/javascripts/services/**/*.coffee",
            "app/assets/javascripts/animations/**/*.coffee",
            "app/assets/javascripts/constants/**/*.coffee",
            "app/assets/javascripts/directives/**/*.coffee",
            "app/assets/javascripts/filters/**/*.coffee",
            "app/assets/javascripts/helpers/**/*.coffee",
            "app/assets/javascripts/controllers/**/*.coffee"
        ],
        watch: "app/assets/javascripts/**/*.coffee",
        dest: pub + '/js'
    }
};
