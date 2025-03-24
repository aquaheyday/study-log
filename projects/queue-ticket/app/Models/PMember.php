<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Laravel\Sanctum\HasApiTokens;


class PMember extends Authenticatable
{
    use HasApiTokens, HasFactory, Notifiable;
    public function getAuthIdentifierName()
    {
        return 'ID';
    }
    public function getAuthPassword()
    {
        return $this->attributes['PASSWORD'];
    }
    protected $fillable = [
        'ID',
        'PASSWORD',
    ];

    protected $hidden = [
        'password',
    ];

    protected $table = 'P_MEMBER';
    public $timestamps = false;

    protected $connection = 'mysql';
}
