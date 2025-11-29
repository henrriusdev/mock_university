<?php

namespace Database\Factories;

use App\Models\Career;
use App\Models\Professor;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

/**
 * @extends Factory<Career>
 */
class CareerFactory extends Factory
{
    protected $model = Career::class;

    public function definition(): array
    {
        return [
            'name' => Str::title(fake()->unique()->words(3, true)),
            'description' => fake()->optional()->paragraph(),
            'code' => strtoupper(fake()->unique()->bothify('CAR-###')),
            'leader_id' => fake()->boolean(60) ? Professor::factory() : null,
        ];
    }
}
