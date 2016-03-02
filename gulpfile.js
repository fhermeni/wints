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
var print = require("gulp-print");

var config = {
    production: !!util.env.production
};

gulp.task('html', function() {
  return gulp.src('assets/html/*.html')
    .pipe(config.production ? htmlmin({collapseWhitespace: true}) : util.noop())
    .pipe(gulp.dest('assets'))
    .pipe(livereload());
});

//gulp.task('fonts', function() {
//  return gulp.src('assets/css/fonts/**/*')
//  .pipe(gulp.dest('dist/css/fonts'))
//  .pipe(livereload());
//})

gulp.task('handlebars', function(){
  gulp.src(['assets/hbs/*.hbs','assets/hbs/*.partial'])
    .pipe(handlebars())
    .pipe(wrap('Handlebars.template(<%= contents %>)'))
    .pipe(declare({
      namespace: 'wints.templates',
      noRedeclare: true, // Avoid duplicate declarations 
    }))
    .pipe(concat('templates.js'))
    .pipe(gulp.dest('assets/js/'))
    .pipe(livereload());
});

gulp.task('js', function(){      
    return gulp.src(['assets/js/**/*.js', '!assets/js/wints.min.js'])
    .pipe(order([
    "**/j*.min.js",
    "/**/bootstrap.min.js",    
    "**/m*.min.js",
    "**/*.min.js",
    "assets/js/users.js"    
  ]))        
    .pipe(print())
        .pipe(config.production ? uglify() : util.noop())        
        .pipe(concat('wints.min.js'))                      
        .pipe(gulp.dest('assets/js/'))        
        .pipe(livereload());        
});

gulp.task('css', function() {
  return gulp.src(['assets/css/**/*.css', '!assets/js/**/*.min.css', '!assets/js/**/wints.css'])
    .pipe(config.production ? cleanCSS() : util.noop())
    .pipe(concat('wints.css'))                      
    .pipe(gulp.dest('assets/css/'))            
    .pipe(livereload());
});

gulp.task('watch', function(){
  livereload.listen();
  gulp.watch(['assets/js/**/*.js', '!assets/js/**/*.min.js', '!assets/js/wints.js'], ['js']);   
  gulp.watch(['assets/css/**/*.css', '!assets/css/**/*.min.css', '!assets/css/wints.css'], ['css']); 
  gulp.watch('assets/html/*.html', ['html']);
  //gulp.watch('assets/css/fonts/**/*', ['fonts']); 
})
