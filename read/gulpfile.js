var gulp = require('gulp'),
    cleanCss = require("gulp-clean-css"),
    uglify = require('gulp-uglify'),
    rename = require('gulp-rename'),
    babel = require("gulp-babel"),
    watch = require('gulp-watch'),
    browserSync = require('browser-sync').create(),
    clean = require('gulp-clean')

gulp.task('browser-reload', (done) => {
    browserSync.reload()
    done()
})

gulp.task('css', () => {
    return gulp.src('public/static/css/*.css')
        .pipe(cleanCss())
        .pipe(rename({suffix: '.min'}))
        .pipe(gulp.dest('public/dist/css'))
});

gulp.task('js', () => {
    return gulp.src(['public/static/js/*.js'])
        .pipe(babel({
            presets: ['es2015']
        }))
        .pipe(uglify({
            mangle: true,
            output: {ascii_only: true}
        }))

        .pipe(rename({suffix: '.min'}))
        .pipe(gulp.dest('public/dist/js'))
});

// clean
gulp.task('clean', () => {
    return gulp.src(['public/dist/*'])
        .pipe(clean())
})

gulp.task('build', gulp.series('clean', gulp.parallel('css', 'js')))

gulp.task('browser', function(done) {
    browserSync.init({
        proxy: "127.0.0.1:8080",
        port: 2000
    })
    done()
})

gulp.task('watch', function() {
    watch([
        "public/static/css/*.css",
        "public/static/js/*.js"
    ], gulp.series('build', 'browser-reload'))
})

gulp.task('serve', gulp.series('build', 'browser', 'watch'))