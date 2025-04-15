<?php

namespace App\Http\Controllers;

use App\Models\PNumber;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class KioskController extends Controller
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
            ->where('STATUS', '01')
            ->get();

        $result = [];

        foreach ($country as $item) {
            if ($item == 'kr') {
                $result[$item] = [
                    'fir_text' => '대기자 ',
                    'number' => $data->where('COUNTRY', $item)->count(),
                    'last_text' => '명',
                    'button' => '접수하기'
                ];
            } else if ($item == 'jp') {
                $result[$item] = [
                    'fir_text' => '待機者 ',
                    'number' => $data->where('COUNTRY', $item)->count(),
                    'last_text' =>  '人',
                    'button' => '受付する'
                ];
            } else if ($item == 'ch') {
                $result[$item] = [
                    'fir_text' => '',
                    'number' => $data->where('COUNTRY', $item)->count(),
                    'last_text' =>  '人等候',
                    'button' => '申請'
                ];
            } else if ($item == 'en') {
                $result[$item] = [
                    'fir_text' => '',
                    'number' => $data->where('COUNTRY', $item)->count(),
                    'last_text' =>  ' people waiting',
                    'button' => 'Apply'
                ];
            }

        }

        return view('kiosk', [
            'data' => $result,
        ]);
    }

    public function call(Request $request)
    {
        $result = [
            'success' => false,
            'text' => '',
            'sub_text' => '',
            'number' => '',
            'button' => ''
        ];

        $model = new PNumber();

        $defaultNo = $model->where('BRANCH', Auth::user()->BRANCH)
            ->where('COUNTRY', $request->input('country'))
            ->where('DATE', date('Y-m-d'))
            ->orderBy('NO', 'DESC')
            ->first();

        $counsel = $model->where('BRANCH', Auth::user()->BRANCH)
            ->where('COUNTRY', $request->input('country'))
            ->where('DATE', date('Y-m-d'))
            ->orderBy('COUNTRY_COUNSEL', 'DESC')
            ->first();

        $no = is_null($defaultNo) || empty($defaultNo->NO) ? 1 : $defaultNo->NO + 1;
        $counselNo = is_null($counsel) || empty($counsel->COUNTRY_COUNSEL) ? 1 : $counsel->COUNTRY_COUNSEL + 1;

        $model->BRANCH = Auth::user()->BRANCH;
        $model->COUNTRY = $request->input('country');
        $model->NO = $no;
        $model->COUNTRY_COUNSEL = $counselNo;
        $model->DATE = date('Y-m-d');
        $model->save();

        if ($request->input('country') == 'kr') {
            $result['text'] = '접수 대기 등록이 완료되었습니다.';
            $result['sub_text'] = '대기번호';
            $result['number'] = $no;
            $result['button'] = '확인';
        } elseif ($request->input('country') == 'jp') {
            $result['text'] = '受付待ち登録が完了しました。';
            $result['sub_text'] = '待ち番号';
            $result['number'] = $no;
            $result['button'] = '確認';
        } elseif ($request->input('country') == 'ch') {
            $result['text'] = '等待登記已完成。';
            $result['sub_text'] = '等候號碼';
            $result['number'] = $no;
            $result['button'] = '查看';
        } elseif ($request->input('country') == 'en') {
            $result['text'] = 'Waiting registration has been completed.';
            $result['sub_text'] = 'Waiting Number';
            $result['number'] = $no;
            $result['button'] = 'Ok';
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
