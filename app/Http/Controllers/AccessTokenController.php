<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;

class AccessTokenController extends Controller
{
    public function get(Request $request) {
        $clientId = $request->header('ClientID');

        if ($clientId == env('CLIENT_ID_IOS')) {
            return response()->json([
                'code' => '0001',
                'token' => env('ACCESS_TOKEN')
            ]);
        }

        return response()->json([
            'code' => '1001'
        ]);
    }
}