<?php

namespace App\Http\Controllers;

use App\Models\Article;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class ArticleController extends Controller
{
    public function get() {
        \Log::debug('get articles.');

        // // ダミーデータ作成
        $article = new Article;
        $article->source_name = "test";
        $article->author = "test";
        $article->title = "test";
        $article->description = "test";
        $article->url = "test";
        $article->image_url = "test";
        $article->published_time = "test";
        $article->save();

        //検索
        $articles = Article::all();
        foreach ($articles as $row) {
            \Log::debug($row->url);
        }

        $path = storage_path() . "/mock/article_data.json";
        $aricle_data = json_decode(file_get_contents($path), true);

        return response()->json([
            'code' => 200,
            'data' => $aricle_data
        ]);
    }

    public function put() {
        \Log::debug('update articles.');
        return response()->json([
            'code' => 200
        ]);
    }
}
