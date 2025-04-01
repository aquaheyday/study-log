<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Notifications\Notifiable;
use Laravel\Sanctum\HasApiTokens;

class PNumber extends Model
{
    protected $table = 'P_NUMBER';
    public $timestamps = false;
    protected $connection = 'mysql';
}
