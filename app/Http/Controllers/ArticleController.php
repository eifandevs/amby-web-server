<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ArticleController extends Controller
{
    public function index() {
        \Log::debug('article get.');
        $path = storage_path() . "/mock/article_data.json";
        $aricle_data = json_decode(file_get_contents($path), true);
        return $aricle_data;
    }
}
