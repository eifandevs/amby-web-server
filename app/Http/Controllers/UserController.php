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
        $id = $request->input('id');
        $password = $request->input('password');
        $email = $request->input('email');
        $user = new User;
        $user->id = $id;
        $user->password = $password;
        $user->email = $email;
        $user->save();
        return "<html>ok</html>";
    }
}