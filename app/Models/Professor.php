<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Professor extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'user_id',
        'boss_id',
    ];

    /**
     * @return BelongsTo<User, self>
     */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class);
    }

    /**
     * @return BelongsTo<self, self>
     */
    public function boss(): BelongsTo
    {
        return $this->belongsTo(self::class, 'boss_id');
    }

    /**
     * @return HasMany<self>
     */
    public function subordinates(): HasMany
    {
        return $this->hasMany(self::class, 'boss_id');
    }

    /**
     * @return HasMany<Career>
     */
    public function ledCareers(): HasMany
    {
        return $this->hasMany(Career::class, 'leader_id');
    }

    /**
     * @return HasMany<Subject>
     */
    public function subjects(): HasMany
    {
        return $this->hasMany(Subject::class);
    }

    /**
     * @return HasMany<Request>
     */
    public function requests(): HasMany
    {
        return $this->hasMany(Request::class);
    }
}
