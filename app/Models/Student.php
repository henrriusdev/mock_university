<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Student extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'credit_units',
        'total_average',
        'semester',
        'user_id',
        'career_id',
    ];

    protected $casts = [
        'credit_units' => 'integer',
        'total_average' => 'float',
        'semester' => 'integer',
    ];

    /**
     * @return BelongsTo<User, self>
     */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class);
    }

    /**
     * @return BelongsTo<Career, self>
     */
    public function career(): BelongsTo
    {
        return $this->belongsTo(Career::class);
    }

    /**
     * @return HasMany<Note>
     */
    public function notes(): HasMany
    {
        return $this->hasMany(Note::class);
    }

    /**
     * @return HasMany<Payment>
     */
    public function payments(): HasMany
    {
        return $this->hasMany(Payment::class);
    }

    /**
     * @return HasMany<Request>
     */
    public function requests(): HasMany
    {
        return $this->hasMany(Request::class);
    }
}
