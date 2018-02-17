<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

\Log::debug('html request.');

Route::get('/', 'IndexController@index');

// Route::get('/', function () {
//     $file = '../';
//     return '<html><body><h1>Hello World</h1><p>this is sample page.</p></body></html>';
//     // return view('welcome');
// });

