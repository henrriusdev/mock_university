<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Role extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'name',
        'description',
    ];

    /**
     * @return BelongsToMany<Permissions>
     */
    public function permissions(): BelongsToMany
    {
        return $this->belongsToMany(Permissions::class, 'roles_permissions');
    }

    /**
     * @return HasMany<User>
     */
    public function users(): HasMany
    {
        return $this->hasMany(User::class);
    }
}
