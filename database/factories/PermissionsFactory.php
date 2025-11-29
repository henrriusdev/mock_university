<?php

namespace Database\Factories;

use App\Models\Permissions;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

/**
 * @extends Factory<Permissions>
 */
class PermissionsFactory extends Factory
{
    protected $model = Permissions::class;

    public function definition(): array
    {
        return [
            'name' => Str::snake(fake()->unique()->words(3, true)),
            'description' => fake()->optional()->sentence(),
        ];
    }
}
