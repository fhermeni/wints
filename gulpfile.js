/*
 * Copyright (c) 2015 Martin Donath
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 */

/* ----------------------------------------------------------------------------
 * Imports
 * ------------------------------------------------------------------------- */

var gulp       = require('gulp');
var addsrc     = require('gulp-add-src');
var args       = require('yargs').argv;
var autoprefix = require('autoprefixer-core');
var clean      = require('del');
var collect    = require('gulp-rev-collector');
var concat     = require('gulp-concat');
var ignore     = require('gulp-ignore');
var mincss     = require('gulp-minify-css');
var minhtml    = require('gulp-htmlmin');
var modernizr  = require('gulp-modernizr');
var mqpacker   = require('css-mqpacker');
var notifier   = require('node-notifier');
var gulpif     = require('gulp-if');
var pixrem     = require('pixrem');
var plumber    = require('gulp-plumber');
var ren    = require('gulp-rename');
var postcss    = require('gulp-postcss');
var reload     = require('gulp-livereload');
var rev        = require('gulp-rev');
var sass       = require('gulp-sass');
var sourcemaps = require('gulp-sourcemaps');
var sync       = require('gulp-sync')(gulp).sync;
var child      = require('child_process');
var uglify     = require('gulp-uglify');
var util       = require('gulp-util');
var vinyl      = require('vinyl-paths');

/* ----------------------------------------------------------------------------
 * Locals
 * ------------------------------------------------------------------------- */

/* Application server */
var server = null;

/* ----------------------------------------------------------------------------
 * Overrides
 * ------------------------------------------------------------------------- */

/*
 * Override gulp.src() for nicer error handling.
 */
var src = gulp.src;
gulp.src = function() {
  return src.apply(gulp, arguments)
    .pipe(plumber(function(error) {
      util.log(util.colors.red(
        'Error (' + error.plugin + '): ' + error.message
      ));
      notifier.notify({
        title: 'Error (' + error.plugin + ')',
        message: error.message.split('\n')[0]
      });
      this.emit('end');
    })
  );
};

/* ----------------------------------------------------------------------------
 * Assets pipeline
 * ------------------------------------------------------------------------- */

var gulp = require('gulp'),
    gp_concat = require('gulp-concat'),
    gp_rename = require('gulp-rename'),
    gp_uglify = require('gulp-uglify');

/*
* Javascript concatenation and minimization
*/
gulp.task('assets:javascripts', function(){
    return gulp.src(['assets/js/*.js','assets/js/vendor/*.js'])
        .pipe(gp_concat('wints.js'))
        .pipe(gulp.dest('dist'))
        .pipe(gp_rename('wints.min.js'))
        .pipe(gp_uglify())
        .pipe(gulp.dest('dist'));
});

/*
* stylesheets concatenation and minimization
*/
gulp.task('assets:stylesheets', function(){
    return gulp.src(['assets/css/*.css','assets/css/vendor/*.css'])
        .pipe(gp_concat('wints.css'))
        .pipe(gulp.dest('dist'))
        .pipe(gp_rename('wints.min.css'))
        .pipe(gp_uglify())
        .pipe(gulp.dest('dist'));
});
/*
 * Build assets.
 */
gulp.task('assets:build', [
  'assets:stylesheets',
  'assets:js',
  'assets:modernizr',
  'assets:views'
]);

/*
 * Watch assets for changes and rebuild on the fly.
 */
gulp.task('assets:watch', function() {

  /* Rebuild stylesheets on-the-fly */
  gulp.watch([
    'assets/css/**/*.css'
  ], ['assets:stylesheets']);

  /* Rebuild javascripts on-the-fly */
  gulp.watch([
    'assets/js/**/*.js',
    'bower.json'
  ], ['assets:javascripts']);
});

/* ----------------------------------------------------------------------------
 * Application server
 * ------------------------------------------------------------------------- */

/*
 * Build assets by default.
 */
gulp.task('default', ['js-fef'], function(){});