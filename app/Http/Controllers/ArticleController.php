<?php

namespace App\Http\Controllers;

use App\Models\Article;
use App\Models\AppSetting;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class ArticleController extends Controller
{
    public function get(Request $request) {
        $token = $request->header('X-Auth-Token');

        if ($token != env('ACCESS_TOKEN')) {
            return response()->json([
                'code' => '0101'
            ]);
        }

        $app_settings = AppSetting::all();
        if (count($app_settings) == 0) {
            \Log::error("no setting.");     
            return response()->json([
                'code' => '9999'
            ]);
        }

        return response()->json([
            'code' => '0001',
            'data' => Article::all(),
            'next' => AppSetting::all()->first()->article_updated_at
        ]);
    }

    /// update articles
    public function put(Request $request) {
        // 取得日の更新
        $app_settings = AppSetting::all();
        if (count(app_settings) > 0) {
            $app_setting = $app_settings->first();
            $app_setting->article_updated_at = date("Y-m-d H:i:s");;
            $app_setting->save();
        } else {
            $new_app_setting = new AppSetting;
            $new_app_setting->article_updated_at = date("Y-m-d H:i:s");
            $new_app_setting->save();
        }

        $token = $request->header('X-Auth-Token');

        if ($token != env('ACCESS_TOKEN')) {
            return response()->json([
                'code' => '0101'
            ]);
        }

        $base_url = 'https://newsapi.org';
        $api_key = env('NEWS_API_KEY');
        $categories = ['business', 'entertainment', 'health', 'science', 'technology', 'sports'];
        $client = new \GuzzleHttp\Client([
            'base_uri' => $base_url,
            'exceptions' => false
        ]);
        $path = '/v2/top-headlines';

        // カテゴリ毎に取得
        foreach ($categories as $category) {
            try {
                $response = $client->request('GET', $path,
                    ['query' => [
                        'country' => 'jp',
                        'category' => $category,
                        'apiKey' => $api_key
                    ]]
                );

                $code = $response->getStatusCode();
                $response_array = json_decode((string) $response->getBody());
            } catch (Exception $e) {
                \Log::error("caught exception. message: $e->getMessage()");     
                return response()->json([
                    'code' => '9999'
                ]);
            }

            $article_index = 0;

            if ($code === 200 && $response_array->status === 'ok') {
                $response_articles = $response_array->articles;
                if (count($response_articles) > 0) {
                    \Log::debug("update $category articles.");

                    // 現在のデータは削除する
                    Article::where('category', "$category")->delete();
                    // サムネイルの削除
                    $thumbnailPath = resource_path() . "assets/image/thumbnails/$category/*";
                    system("rm -rf $thumbnailPath");

                    foreach ($response_articles as $response_article) {
                        $article_index++;

                        // 画像はローカルに保存する
                        if (!is_null($response_article->urlToImage)) {
                            \Log::debug("save image. url: $image_url");
                            try {
                                $image_url = $response_article->urlToImage;
                                $fileName = substr(strrchr($image_url, "/"), 1);
                                $savePath = resource_path() . "/assets/image/thumbnails/$category/$article_index/$fileName";

                                $client = new \GuzzleHttp\Client();
                                $client->get($image_url,[
                                    'save_to' => $savePath,
                                ]);
                            } catch (Exception $e) {
                                \Log::error("caught exception. message: $e->getMessage()");     
                                return response()->json([
                                    'code' => '9999'
                                ]);
                            }
                        }

                        // 記事をDBに保存する
                        $new_article = new Article;
                        $new_article->category = $category;
                        if (!is_null($response_article->source) && !is_null($response_article->source->name)) {
                            $new_article->source_name = $response_article->source->name;
                        }
                        if (!is_null($response_article->author)) {
                            $new_article->author = $response_article->author;
                        }
                        if (!is_null($response_article->title)) {
                            $new_article->title = $response_article->title;
                        }
                        if (!is_null($response_article->description)) {
                            $new_article->description = $response_article->description;
                        }
                        if (!is_null($response_article->url)) {
                            $new_article->url = $response_article->url;
                        }
                        if (!is_null($response_article->urlToImage)) {
                            $new_article->image_url = $response_article->urlToImage;
                        }
                        if (!is_null($response_article->publishedAt)) { 
                            $new_article->published_at = $response_article->publishedAt;
                        }
                        $new_article->save();
                    }
                } else {
                    \Log::error("get no articles.");
                }
            } else {
                \Log::error("failed to update articles. category: $category, httpCode: $code, code: $response_array->code, message: $response_array->message");     
                return response()->json([
                    'code' => '1002'
                ]);           
            }
        }

        return response()->json([
            'code' => '0001'
        ]);
    }
}