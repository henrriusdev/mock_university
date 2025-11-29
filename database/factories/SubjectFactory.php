<?php

namespace Database\Factories;

use App\Models\Career;
use App\Models\Professor;
use App\Models\Subject;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

/**
 * @extends Factory<Subject>
 */
class SubjectFactory extends Factory
{
    protected $model = Subject::class;

    public function definition(): array
    {
        $practiceHours = fake()->numberBetween(0, 4) * 2;
        $theoryHours = fake()->numberBetween(1, 4) * 2;
        $labHours = fake()->numberBetween(0, 4) * 2;
        $totalHours = max($practiceHours + $theoryHours + $labHours, 2);

        $startTime = fake()->time('H:i');
        $endTime = date('H:i', strtotime($startTime) + 2 * 3600);

        return [
            'name' => Str::title(fake()->unique()->words(3, true)),
            'description' => fake()->optional()->paragraph(),
            'credit_units' => fake()->numberBetween(1, 6),
            'semester' => fake()->numberBetween(1, 12),
            'code' => strtoupper(fake()->unique()->bothify('SUBJ###')),
            'practice_hours' => $practiceHours,
            'theory_hours' => $theoryHours,
            'lab_hours' => $labHours,
            'total_hours' => $totalHours,
            'class_schedule' => [
                [
                    'day' => fake()->dayOfWeek(),
                    'start' => $startTime,
                    'end' => $endTime,
                ],
            ],
            'professor_id' => Professor::factory(),
            'career_id' => Career::factory(),
        ];
    }
}
