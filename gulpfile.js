var gulp = require('gulp');
var handlebars = require('gulp-handlebars');
var wrap = require('gulp-wrap');
var declare = require('gulp-declare');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');
var cleanCSS = require('gulp-clean-css');
var htmlmin = require('gulp-htmlmin');
var livereload = require('gulp-livereload');
var util = require('gulp-util');
var order = require("gulp-order");
var merge = require('merge-stream');
var path = require('path');
var print = require('gulp-print');
var config = {
    production: !!util.env.production
};

gulp.task('html', function() {
  return gulp.src('assets/html/*.html')
    .pipe(config.production ? htmlmin({collapseWhitespace: true}) : util.noop())
    .pipe(gulp.dest('assets'))
    .pipe(livereload());
});


gulp.task('templates', function() {
    return gulp.src('assets/hbs/*.partial')
    .pipe(handlebars({
      handlebars: require('handlebars')
    }))
    .pipe(wrap('Handlebars.registerPartial(<%= processPartialName(file.relative) %>, Handlebars.template(<%= contents %>));', {}, {
      imports: {
        processPartialName: function(fileName) {
          return JSON.stringify(path.basename(fileName, '.js'));
        }
      }
    }))
    .pipe(concat('hbs_partial.js'))
    .pipe(gulp.dest('assets/js/'))
    ;
});

gulp.task('partials', function() {
    return gulp.src('assets/hbs/*.hbs')
    .pipe(handlebars({
      handlebars: require('handlebars')
    }))
    .pipe(wrap('Handlebars.template(<%= contents %>)'))
    .pipe(declare({
      namespace: 'wints.templates',
      noRedeclare: true
    }))
    .pipe(concat('hbs.js'))
    .pipe(gulp.dest('assets/js/'));
});

gulp.task('js', function(){
    return gulp.src(['assets/js/**/*.js', '!assets/js/wints.min.js'])
    .pipe(order([
    "**/vendor/jquery-1*",
    "**/vendor/moment.min.js",
    "**/vendor/moment-timezone-with-data-2010-2020.min.js",
    "**/vendor/handlebars.min.js",
    "**/vendor/jquery.*",
    "**/vendor/bootstrap.min.js",
    "**/vendor/bootstrap-*.js",
    "**/vendor/*.js",
    "assets/js/*.js"
  ]))
        .pipe(config.production ? uglify() : util.noop())
        .pipe(concat('wints.min.js'))
        .pipe(gulp.dest('assets/js/'))
        .pipe(livereload());
});

gulp.task('css', function() {
  return gulp.src(['assets/css/**/*.css', '!assets/css/**/wints.css'])
  .pipe(order([
    "**/vendor/*.css",
    "assets/css/*.css"
  ]))
    .pipe(config.production ? cleanCSS() : util.noop())
    .pipe(concat('wints.css'))
    .pipe(gulp.dest('assets/css/'))
    .pipe(livereload());
});

gulp.task('assets', function () {
  return gulp.start(['css','html','partials','templates','js']);
});

gulp.task('watch', function(){
  livereload.listen();
  gulp.watch(['assets/js/**/*.js', '!assets/js/**/*.min.js', '!assets/js/wints.js', '!assets/js/!hbs*.js'], ['js']);
  gulp.watch(['assets/css/**/*.css', '!assets/css/**/*.min.css', '!assets/css/wints.css'], ['css']);
  gulp.watch('assets/html/*.html', ['html']);
  gulp.watch('assets/hbs/*', ['partials','templates']);
})
