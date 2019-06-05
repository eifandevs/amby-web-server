<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Log;
use Illuminate\Http\Request;
use App\User;

class UserController extends Controller
{
    public function get() {
        return "<html>user list</html>";
    }

    public function post(Request $request) {
        $user_id = $request->input('user_id');
        $user = new User;
        $user->user_id = $request->user_id;
        $user->save();
        return "<html>ok</html>";
    }
}