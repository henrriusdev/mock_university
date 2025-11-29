<?php

namespace Database\Factories;

use App\Models\Professor;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends Factory<Professor>
 */
class ProfessorFactory extends Factory
{
    protected $model = Professor::class;

    public function definition(): array
    {
        return [
            'user_id' => User::factory(),
            'boss_id' => null,
        ];
    }

    /**
     * Indicate that the professor reports to another professor.
     */
    public function withBoss(Professor $boss = null): static
    {
        return $this->state(function () use ($boss) {
            return [
                'boss_id' => $boss?->id ?? Professor::factory(),
            ];
        });
    }
}
