<?php

namespace App\Http\Controllers;

use App\Models\Article;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class ArticleController extends Controller
{
    public function get() {
        //検索
        $articles = Article::all();
        foreach ($articles as $row) {
            \Log::debug($row->url);
        }

        $path = storage_path() . "/mock/article_data.json";
        $aricle_data = json_decode(file_get_contents($path), true);

        return response()->json([
            'code' => '0000',
            'data' => $aricle_data
        ]);
    }

    /// 記事更新
    public function put() {
        $base_url = 'https://newsapi.org';
        $api_key = 'baab6473c743412394dc823092be475a';
        $categories = ['business'];
        $client = new \GuzzleHttp\Client( [
            'base_uri' => $base_url,
        ]);
        $path = '/v2/top-headlines';

        // every category request
        foreach ($categories as $category) {
            $response = $client->request('GET', $path,
                ['query' => [
                    'country' => 'jp',
                    'category' => $category,
                    'apiKey' => $api_key
                ]]
            );
            $code = $response->getStatusCode();
            $response_array = json_decode((string) $response->getBody());

            if ($code === 200 && $response_array->status === 'ok') {
                $response_articles = $response_array->articles;
                if (count($response_articles) > 0) {
                    \Log::debug("update $category articles.");

                    Article::truncate();

                    foreach ($response_articles as $response_article) {
                        $new_article = new Article;
                        $new_article->source_name = $response_article->source->name;
                        $new_article->author = $response_article->author;
                        $new_article->title = $response_article->title;
                        $new_article->description = $response_article->description;
                        $new_article->url = $response_article->url;
                        $new_article->image_url = $response_article->urlToImage;
                        $new_article->published_time = $response_article->publishedAt;
                        $new_article->save();
                    }
                }
            } else {
                \Log::error('failed to update $category articles.');                
            }
        }

        return response()->json([
            'code' => 'OK'
        ]);
    }
}
