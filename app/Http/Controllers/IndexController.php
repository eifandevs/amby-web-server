<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class IndexController extends Controller
{
    public function index() {
        $path = storage_path() . "/index.html";
        $client_html = file_get_contents($path, true);
        return $client_html;
    }
}
