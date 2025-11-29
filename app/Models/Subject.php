<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Subject extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'name',
        'description',
        'credit_units',
        'semester',
        'code',
        'practice_hours',
        'theory_hours',
        'lab_hours',
        'total_hours',
        'class_schedule',
        'professor_id',
        'career_id',
    ];

    protected $casts = [
        'credit_units' => 'integer',
        'semester' => 'integer',
        'practice_hours' => 'integer',
        'theory_hours' => 'integer',
        'lab_hours' => 'integer',
        'total_hours' => 'integer',
        'class_schedule' => 'array',
    ];

    /**
     * @return BelongsTo<Professor, self>
     */
    public function professor(): BelongsTo
    {
        return $this->belongsTo(Professor::class);
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
     * @return BelongsToMany<Subject>
     */
    public function prerequisites(): BelongsToMany
    {
        return $this->belongsToMany(self::class, 'precedence', 'subject_id', 'precedence_id');
    }

    /**
     * @return BelongsToMany<Subject>
     */
    public function unlocks(): BelongsToMany
    {
        return $this->belongsToMany(self::class, 'precedence', 'precedence_id', 'subject_id');
    }
}
