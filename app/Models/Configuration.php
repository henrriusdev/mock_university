<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Configuration extends Model
{
    use HasFactory, HasUuids;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $fillable = [
        'start_registration_subjects',
        'end_registration_subjects',
        'block_not_pay_inscription',
        'fee_dates',
        'number_fees',
        'number_notes',
        'notes_percentages',
        'cycle_id',
    ];

    protected $casts = [
        'start_registration_subjects' => 'datetime',
        'end_registration_subjects' => 'datetime',
        'block_not_pay_inscription' => 'boolean',
        'fee_dates' => 'array',
        'number_fees' => 'integer',
        'number_notes' => 'integer',
        'notes_percentages' => 'array',
    ];

    /**
     * @return BelongsTo<Cycle, self>
     */
    public function cycle(): BelongsTo
    {
        return $this->belongsTo(Cycle::class);
    }
}
