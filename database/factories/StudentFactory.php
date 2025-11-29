<?php

namespace Database\Factories;

use App\Models\Career;
use App\Models\Student;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends Factory<Student>
 */
class StudentFactory extends Factory
{
    protected $model = Student::class;

    public function definition(): array
    {
        return [
            'credit_units' => fake()->numberBetween(0, 180),
            'total_average' => fake()->randomFloat(2, 0, 20),
            'semester' => fake()->numberBetween(1, 12),
            'user_id' => User::factory(),
            'career_id' => Career::factory(),
        ];
    }
}
