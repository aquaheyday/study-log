<?php

namespace App\Http\Controllers;

use App\Models\PNumber;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class CallController extends Controller
{
    public function index()
    {
        $country = [
            'kr',
            'jp',
            'ch',
            'en'
        ];

        return view('call', [
            'country' => $country,
        ]);
    }

    public function list(Request $request)
    {
        $result = [
            'success' => false,
            'data' => []
        ];

        $model = new PNumber();

        $country = [
            'kr',
            'jp',
            'ch',
            'en'
        ];

        foreach ($country as $item) {
            $result['data'][$item] = [];

            $data = $model->where('BRANCH', Auth::user()->BRANCH)
                ->where('DATE', date('Y-m-d'))
                ->where('STATUS', '02')
                ->where('COUNTRY', $item)
                ->orderBy('COUNTRY_COUNSEL', 'DESC')
                ->limit(4)
                ->get();

            foreach ($data as $item2) {
                if ($item == 'kr') {
                    $result['data'][$item][] = [
                        'number' => $item2->NO,
                        'counsel' => $item2->ADMIN_NO . ' 상담 창구'
                    ];
                } else if ($item == 'jp') {
                    $result['data'][$item][] = [
                        'number' => $item2->NO,
                        'counsel' => $item2->ADMIN_NO . ' 相談窓口'
                    ];
                } else if ($item == 'ch') {
                    $result['data'][$item][] = [
                        'number' => $item2->NO,
                        'counsel' => $item2->ADMIN_NO . ' 諮詢台'
                    ];
                } else if ($item == 'en') {
                    $result['data'][$item][] = [
                        'number' => $item2->NO,
                        'counsel' => $item2->ADMIN_NO . ' Consultation Desk'
                    ];
                }
            }
        }

        $result['success'] = true;

        return response()->json($result);
    }

    public function waiting()
    {
        $result = [
            'success' => false
        ];

        $model = new PNumber();

        $country = [
            'kr',
            'jp',
            'ch',
            'en'
        ];

        $data = $model->where('BRANCH', Auth::user()->BRANCH)
            ->where('DATE', date('Y-m-d'))
            ->where('STATUS', '01')
            ->get();

        foreach ($country as $item) {
            $result['data'][$item] = $data->where('COUNTRY', $item)->count();
        }

        $result['success'] = true;

        return response()->json($result);
    }
}
