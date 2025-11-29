<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Career extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'name',
        'description',
        'code',
        'leader_id',
    ];

    /**
     * @return BelongsTo<Professor, self>
     */
    public function leader(): BelongsTo
    {
        return $this->belongsTo(Professor::class, 'leader_id');
    }

    /**
     * @return HasMany<Student>
     */
    public function students(): HasMany
    {
        return $this->hasMany(Student::class);
    }

    /**
     * @return HasMany<Subject>
     */
    public function subjects(): HasMany
    {
        return $this->hasMany(Subject::class);
    }
}
