<?php

namespace App\Http\Controllers;

use App\Article;
use Illuminate\Http\Request;

class ArticleController extends Controller
{
    public function index() {
        \Log::debug('article get.');

        // ダミーデータ作成
        // $article = new Article;
        // $article->source_name = "test";
        // $article->author = "test";
        // $article->title = "test";
        // $article->description = "test";
        // $article->url = "test";
        // $article->image_url = "test";
        // $article->published_time = "test";
        // $article->save();

        // 検索
        // $articles = Article::all();
        // foreach ($articles as $row) {
        //     \Log::debug($row->url);
        // }

        $path = storage_path() . "/mock/article_data.json";
        $aricle_data = json_decode(file_get_contents($path), true);
        return $aricle_data;
    }
}
