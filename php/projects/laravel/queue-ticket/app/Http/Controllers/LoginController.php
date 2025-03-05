<?php

namespace App\Http\Controllers;

use App\Models\PMember;
use Illuminate\Support\Facades\Validator;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class LoginController extends Controller
{
    public function index()
    {
        return view('login');
    }

    public function login(Request $request)
    {
        $result = [
            'success' => false,
            'account_type' => ''
        ];

        $validator = Validator::make($request->all(), [
            'id' => 'required|string',
            'password' => 'required|string',
        ]);

        if (!$validator->fails()) {
            $model = new PMember();

            $user = $model->where('ID', $request->input('id'))
                ->where('PASSWORD', $request->input('password'))
                ->first();

            if (!is_null($user)) {
                if ($user->ACCOUNT_TYPE == '01') {
                    Auth::guard('web')->login($user);
                    $result['success'] = true;
                    $result['account_type'] = $user->ACCOUNT_TYPE;
                } elseif ($user->ACCOUNT_TYPE == '02') {
                    Auth::guard('kiosk')->login($user);
                    $result['success'] = true;
                    $result['account_type'] = $user->ACCOUNT_TYPE;
                } elseif ($user->ACCOUNT_TYPE == '03') {
                    Auth::guard('call')->login($user);
                    $result['success'] = true;
                    $result['account_type'] = $user->ACCOUNT_TYPE;
                }
            }
        }

        return response()->json($result);
    }

    public function logout()
    {
        auth::guard('web')->logout();
        auth::guard('kiosk')->logout();
        auth::guard('call')->logout();

        return redirect()->route('login');
    }
}
