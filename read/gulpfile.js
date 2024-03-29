var gulp = require('gulp'),
    cleanCss = require("gulp-clean-css"),
    uglify = require('gulp-uglify'),
    rename = require('gulp-rename'),
    babel = require("gulp-babel"),
    watch = require('gulp-watch'),
    concat = require('gulp-concat'),
    rev = require('gulp-rev'),
    revCollector = require('gulp-rev-collector'),
    browserSync = require('browser-sync').create(),
    clean = require('gulp-clean')

gulp.task('browser-reload', (done) => {
    browserSync.reload()
    done()
})

gulp.task('css', () => {
    return gulp.src('public/static/css/*.css')
        .pipe(cleanCss())
        // .pipe(rename({suffix: '.min'}))
        // .pipe(gulp.dest('public/dist/css'))
        .pipe(concat('app.min.css'))
        .pipe(gulp.dest('public/dist/css'))
        .pipe(rev())
        .pipe(gulp.dest('public/dist/css'))
        .pipe(rev.manifest())
        .pipe(gulp.dest('public/dist/css'))
});

gulp.task('js', () => {
    return gulp.src(['public/static/js/main.js', 'public/static/js/*.js'])
        .pipe(babel({
            presets: ['es2015']
        }))
        .pipe(uglify({
            mangle: true,
            output: {ascii_only: true}
        }))
        // .pipe(rename({suffix: '.min'}))
        // .pipe(gulp.dest('public/dist/js'))
        .pipe(concat('app.min.js'))
        .pipe(gulp.dest('public/dist/js'))
        .pipe(rev())
        .pipe(gulp.dest('public/dist/js'))
        .pipe(rev.manifest())
        .pipe(gulp.dest('public/dist/js'))
});

gulp.task('rev-html', function () {
    return gulp.src(['public/dist/*/*.json', 'app/view/*.php'])
        .pipe(revCollector({
            replaceReved: true
        }))
        .pipe(gulp.dest('app/view/dist'));
});

// clean
gulp.task('clean', (done) => {
    // return gulp.src(['public/dist/*'])
    //     .pipe(clean())
    done()
})

gulp.task('markdown', () => {
    // return gulp.src(['node_modules/moonprism-markdown/markdown.js'])
    //     .pipe(babel({
    //         presets: ['es2015']
    //     }))
    //     .pipe(uglify({
    //         mangle: true,
    //         output: {ascii_only: true}
    //     }))
    //     .pipe(rename({suffix: '.min'}))
    //     .pipe(gulp.dest('public/dist/js'))
    return gulp.src('node_modules/moonprism-markdown/dist/markdown.min.js')
        .pipe(gulp.dest('public/dist/js'))
})

gulp.task('build', gulp.series('clean', gulp.series(gulp.parallel( gulp.parallel('css', 'js'), 'markdown'), 'rev-html')))

gulp.task('default', gulp.series('build'))

gulp.task('browser', function(done) {
    browserSync.init({
        proxy: "127.0.0.1:8023",
        port: 8001
    })
    done()
})

gulp.task('watch', function() {
    watch([
        "public/static/css/*.css",
        "public/static/js/*.js",
        "app/view/*.php",
    ], gulp.series('build', 'browser-reload'))
    // watch([
    //      "app/view/*/*.php",
    // ], gulp.series('browser-reload'))
})

gulp.task('serve', gulp.series('build', 'browser', 'watch'))
