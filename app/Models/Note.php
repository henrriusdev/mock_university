<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Note extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'notes',
        'average',
        'student_id',
        'subject_id',
        'cycle_id',
    ];

    protected $casts = [
        'notes' => 'array',
        'average' => 'float',
    ];

    /**
     * @return BelongsTo<Student, self>
     */
    public function student(): BelongsTo
    {
        return $this->belongsTo(Student::class);
    }

    /**
     * @return BelongsTo<Subject, self>
     */
    public function subject(): BelongsTo
    {
        return $this->belongsTo(Subject::class);
    }

    /**
     * @return BelongsTo<Cycle, self>
     */
    public function cycle(): BelongsTo
    {
        return $this->belongsTo(Cycle::class);
    }
}
