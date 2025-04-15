@extends('layouts.app')

@section('title', '로그인')

@section('styles')
    <link rel="stylesheet" href="/css/login.css">
@endsection

@section('content')
    <div class="login-wrapper">
        <form id="frm">
            <input type="text" name="id" placeholder="id">
            <input type="password" name="pw" placeholder="pw">
            <button type="button" id="loginBtn">로그인</button>
        </form>
    </div>
@endsection

@section('scripts')
    <script src="/js/login.js"></script>
@endsection
