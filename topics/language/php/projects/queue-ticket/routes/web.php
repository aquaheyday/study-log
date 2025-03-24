<?php

use App\Http\Controllers\KioskController;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\LoginController;
use App\Http\Controllers\MainController;
use App\Http\Controllers\CallController;

/*Route::get('/', function (){
    return view('welcome');
})->name('home');*/

Route::get('/', [LoginController::class, 'index'])->name('login');
Route::post('/login', [LoginController::class, 'login']);
Route::get('/logout', [LoginController::class, 'logout'])->name('logout');

Route::middleware(['auth:web'])->group(function () {
    Route::group(['prefix' => 'main'], function () {
        Route::get('/', [MainController::class, 'index'])->name('main');
        Route::post('/first', [MainController::class, 'first']);
        Route::post('/second', [MainController::class, 'second']);
        Route::post('/end', [MainController::class, 'end']);
        Route::get('/list', [MainController::class, 'list']);
    });
});

Route::middleware(['auth:kiosk'])->group(function () {
    Route::group(['prefix' => 'kiosk'], function () {
       Route::get('/', [KioskController::class, 'index'])->name('kiosk');
       Route::post('/call', [KioskController::class, 'call']);
       Route::get('/waiting', [KioskController::class, 'waiting']);
    });
});

Route::middleware(['auth:call'])->group(function () {
    Route::group(['prefix' => 'call'], function () {
        Route::get('/', [CallController::class, 'index'])->name('call');
        Route::get('/list', [CallController::class, 'list']);
    });
});
