<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;

class AccessTokenController extends Controller
{
    public function get(Request $request) {
        $serviceName = $request->header('ServiceName');

        if ($serviceName == 'qass-ios') {
            return response()->json([
                'code' => '0001',
                'token' => env('ACCESS_TOKEN')
            ]);
        }

        return response()->json([
            'code' => '0100'
        ]);
    }
}