<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class UserController extends Controller
{
    public function get() {
        return "<html>user list</html>";
    }

    public function post(Request $request) {
        return "<html>user post</html>";
    }
}