<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Request extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'title',
        'description',
        'status',
        'student_id',
        'professor_id',
    ];

    /**
     * @return BelongsTo<Student, self>
     */
    public function student(): BelongsTo
    {
        return $this->belongsTo(Student::class);
    }

    /**
     * @return BelongsTo<Professor, self>
     */
    public function professor(): BelongsTo
    {
        return $this->belongsTo(Professor::class);
    }
}
