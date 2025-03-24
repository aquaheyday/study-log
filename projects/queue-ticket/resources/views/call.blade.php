@extends('layouts.app')

@section('title', '대기자 호출 화면')

@section('styles')
    <link rel="stylesheet" href="/css/kiosk.css">
@endsection

@section('content')
    <div>
        @foreach($country as $item)
            <div>
                <span>{{ $item }}</span>
                <div class="{{ $item }}_list"></div>
            </div>
        @endforeach
    </div>
@endsection

@section('scripts')
    <script src="/js/call.js"></script>
@endsection
