<?php

use App\Http\Controllers\CareerController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::middleware('auth')->group(function () {
    Route::resource('careers', CareerController::class)
        ->except(['index', 'show']);
});
