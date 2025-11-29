<?php

namespace Database\Factories;

use App\Models\Cycle;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Carbon;

/**
 * @extends Factory<Cycle>
 */
class CycleFactory extends Factory
{
    protected $model = Cycle::class;

    public function definition(): array
    {
        $startDate = Carbon::parse(fake()->dateTimeBetween('-1 year', '+1 year'));
        $endDate = (clone $startDate)->addMonths(fake()->numberBetween(3, 6));

        return [
            'name' => 'Cycle ' . $startDate->format('Y-m'),
            'start_date' => $startDate,
            'end_date' => $endDate,
            'is_active' => fake()->boolean(30),
        ];
    }
}
