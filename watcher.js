#!/usr/bin/env node

// dpw@alameda.local
// 2015.03.04
'use strict';

var fs = require('fs'),
    spawn = require('child_process').spawn,
    clearScreen = '[H[2J',
    files = [],
    tid,
    lastRun;

var run = function() {
    process.stdout.write( clearScreen ); 

    try {
        var cmd = 'make';
        var runner = spawn( cmd, [ 'testmodels' ] );

        runner.stdout.on('data', function( data ) {
            process.stdout.write( data );
        });

        runner.stderr.on('data', function( data ) {
            process.stdout.write( data );
        });

        runner.on('close', function(code) {
            if (code !== 0) {
                console.log( cmd, ' did not exit correctly, code: ', code);
            }

            console.log( '------------------------------------ last run: ', new Date().toISOString() );

            tid = null;
            files.splice( 0 );
        });
    } catch (err) {
        console.log( err );
    }
};

var changeHandler = function(event, filename) {
    if (filename.indexOf( '.go') > 0) {
        console.log( 'file change: ', filename);

        files.push( filename );

        if (!tid) {
            tid = setTimeout(function() {
                run();
            }, 500);
        }
    }
};

fs.watch( './src/', { recursive:true }, changeHandler );
fs.watch( './test/', { recursive:true }, changeHandler );

process.stdout.write( clearScreen ); 
console.log('watching go files...');

