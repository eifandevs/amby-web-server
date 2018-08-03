<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

// DB構築
Route::put('/migrate', function () {
    Artisan::call('migrate', [
        '--force' => false,
    ]);
    return 'OK';
});

// アクセストークンの取得
Route::get('/access_token', 'AccessTokenController@get');

// 記事の取得
Route::get('/articles', 'ArticleController@get');

// 記事の更新
Route::put('/articles', 'ArticleController@put');