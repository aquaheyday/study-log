<?php

namespace App\Http\Controllers;

use App\Models\PNumber;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class MainController extends Controller
{
    public function index()
    {
        $model = new PNumber();

        $country = [
            'kr',
            'jp',
            'ch',
            'en'
        ];

        $data = $model->where('BRANCH', Auth::user()->BRANCH)
            ->where('DATE', date('Y-m-d'))
            ->orderBy('COUNTRY', 'DESC')
            ->orderBy('NO', 'DESC')
            ->get();

        $count = $data->where('STATUS', '02')->where('ADMIN_ID', Auth::user()->ID)->count();

        return view('main', [
            'country' => $country,
            'data' => $data,
            'count' => $count
        ]);
    }

    public function first(Request $request)
    {
        $result = [
            'success' => false,
        ];

        $model = new PNumber();

        $check = $model->where('BRANCH', $request->input('branch'))
            ->where('COUNTRY', $request->input('country'))
            ->where('NO', $request->input('no'))
            ->where('DATE', date('Y-m-d'))
            ->first();

        if ($check->STATUS == '01') {
            $counsel = $model->where('BRANCH', $request->input('branch'))
                ->where('COUNTRY', $request->input('country'))
                ->where('DATE', date('Y-m-d'))
                ->orderBy('COUNTRY_COUNSEL', 'DESC')
                ->first();

            $counselNo = is_null($counsel) || empty($counsel->COUNTRY_COUNSEL) ? 1 : $counsel->COUNTRY_COUNSEL + 1;

            $model->where('BRANCH', $request->input('branch'))
                ->where('COUNTRY', $request->input('country'))
                ->where('NO', $request->input('no'))
                ->where('DATE', date('Y-m-d'))
                ->update([
                    'STATUS' => '02',
                    'COUNTRY_COUNSEL' => $counselNo,
                    'ADMIN_ID' => Auth::user()->ID,
                    'ADMIN_NO' => Auth::user()->NO
                ]);

            $result['success'] = true;
        } else {
            $result['admin_no'] = $check->ADMIN_NO;
        }

        return response()->json($result);
    }

    public function second(Request $request)
    {
        $result = [
            'success' => false,
        ];

        $model = new PNumber();

        $counsel = $model->where('BRANCH', $request->input('branch'))
            ->where('COUNTRY', $request->input('country'))
            ->where('DATE', date('Y-m-d'))
            ->orderBy('COUNTRY_COUNSEL', 'DESC')
            ->first();

        $counselNo = $counsel->COUNTRY_COUNSEL + 1;

        $model->where('BRANCH', $request->input('branch'))
            ->where('COUNTRY', $request->input('country'))
            ->where('NO', $request->input('no'))
            ->where('DATE', date('Y-m-d'))
            ->update([
                'COUNTRY_COUNSEL' => $counselNo,
            ]);

        $result['success'] = true;

        return response()->json($result);
    }

    public function end(Request $request)
    {
        $model = new PNumber();

        $model->where('BRANCH', $request->input('branch'))
            ->where('COUNTRY', $request->input('country'))
            ->where('NO', $request->input('no'))
            ->where('DATE', date('Y-m-d'))
            ->update([
               'STATUS' => '03',
                'ADMIN_ID' => 'test3',
                'END_TIME' => date('H:i:s')
            ]);

        $result = [
            'success' => true,
        ];

        return response()->json($result);
    }

    public function list()
    {
        $model = new PNumber();

        $country = [
            'kr',
            'jp',
            'ch',
            'en'
        ];

        $data = [];

        foreach ($country as $key => $value) {
            $data[$value] = $model->where('BRANCH', Auth::user()->BRANCH)
                ->where('DATE', date('Y-m-d'))
                ->where('COUNTRY', $value)
                ->orderBy('COUNTRY', 'DESC')
                ->orderBy('NO', 'DESC')
                ->get();
        }


        $count = $model->where('BRANCH', Auth::user()->BRANCH)
            ->where('DATE', date('Y-m-d'))
            ->where('STATUS', '02')
            ->where('ADMIN_ID', Auth::user()->ID)
            ->count();

        $result = [
            'success' => true,
            'country' => $country,
            'count' => $count,
            'data' => $data
        ];

        return response()->json($result);
    }
}
