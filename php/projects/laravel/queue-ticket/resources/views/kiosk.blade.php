@extends('layouts.app')

@section('title', '키오스크 대기표 화면')

@section('styles')
    <link rel="stylesheet" href="/css/kiosk.css">
@endsection

@section('content')
    <div class="kiosk-container">
        @foreach($data as $key => $item)
            <div class="kiosk-item">
                <p>{{ $item['fir_text'] }}<span id="{{ $key }}_waiting_number">{{ $item['number'] }}</span>{{ $item['last_text'] }}</p>
                <button type="button" class="callBtn" data-country="{{ $key }}">{{ $item['button'] }}</button>
            </div>
        @endforeach

        <div id="customAlert" class="custom-alert hidden">
            <div class="custom-alert-content">
                <p id="alertMessage"></p>
                <button type="button" id="alertOkBtn">OK</button>
            </div>
        </div>
    </div>
@endsection

@section('scripts')
    <script src="/js/kiosk.js"></script>
@endsection
