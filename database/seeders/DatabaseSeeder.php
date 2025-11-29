<?php

namespace Database\Seeders;

use App\Models\Role;
use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        Role::factory()->createMany([
            ['name' => 'Admin', 'description' => 'Administrator role with full permissions'],
            ['name' => 'Professor', 'description' => 'Professor role with limited permissions'],
            ['name' => 'Student', 'description' => 'Student role with basic permissions'],
            ['name' => 'Professor Boss', 'description' => 'Professor Boss role with elevated permissions'],
            ['name' => 'Career Lead', 'description' => 'Career Lead role with specific permissions'],
            ['name' => 'Registrar', 'description' => 'Registrar role with management permissions'],
        ]);

        $adminRole = Role::where('name', 'Admin')->first();

        User::factory()->create([
            'name' => 'Admin User',
            'email' => 'admin@example.com',
            'first_name' => 'Admin',
            'last_name' => 'User',
            'identification_number' => 'A001',
            'dob' => '1980-01-01',
            'phone' => '1234567890',
            'address' => [
                'street' => 'Av. Universidad',
                'city' => 'Caracas',
                'state' => 'Distrito Capital',
                'zip_code' => '1010',
                'country' => 'Venezuela'
            ],
            'password' => bcrypt('password'),
            'profile_picture' => null,
            'is_active' => true,
            'role_id' => $adminRole->id,
        ]);
    }
}
