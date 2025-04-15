@extends('layouts.app')

@section('title', '대기자 관리')

@section('styles')
    <link rel="stylesheet" href="/css/main.css">
@endsection

@section('content')
    <div class="top-buttons">
        <button type="button" onclick="location.reload();">새로고침</button>
        <button type="button" onclick="location.href='{{ route('logout') }}';">로그아웃</button>
    </div>

    <div id="target">
        @foreach($country as $key => $value)
            <div class="country-section">
                <p>{{ $value }}</p>
                @foreach($data->where('COUNTRY', $value) as $key2 => $item)
                    <div class="ticket-item">
                        <span>{{ $item['NO'] }}</span>
                        <span>
                        @if($item['STATUS'] == '01')
                                호출가능
                            @elseif($item['STATUS'] == '02')
                                호출 중
                            @else
                                상담 완료
                            @endif
                    </span>
                        @if($item['STATUS'] == '01')
                            <button type="button" class="firstBtn" @if($count > 0) disabled @endif data-branch="{{ $item['BRANCH'] }}" data-country="{{ $item['COUNTRY'] }}" data-no="{{ $item['NO'] }}">호출 하기</button>
                        @elseif($item['STATUS'] == '02')
                            @if($item['ADMIN_ID'] == Auth::user()->ID)
                                <button type="button" class="secondBtn" data-branch="{{ $item['BRANCH'] }}" data-country="{{ $item['COUNTRY'] }}" data-no="{{ $item['NO'] }}">재 호출</button>
                                <button type="button" class="endBtn" data-branch="{{ $item['BRANCH'] }}" data-country="{{ $item['COUNTRY'] }}" data-no="{{ $item['NO'] }}">상담 완료</button>
                            @else
                                <span>{{ $item['ADMIN_NO'] }} 상담</span>
                                <span>대기중</span>
                            @endif
                        @else
                            <span>{{ $item['ADMIN_NO'] }} 상담</span>
                            <span>{{ $item['END_TIME'] }}</span>
                        @endif
                    </div>
                @endforeach
            </div>
        @endforeach
    </div>
@endsection

@section('scripts')
    <script src="/js/main.js"></script>
@endsection
