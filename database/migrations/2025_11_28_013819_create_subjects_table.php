<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('subjects', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('name');
            $table->text('description')->nullable();
            $table->integer('credit_units');
            $table->integer('semester');
            $table->string('code')->unique();

            // Horas
            $table->integer('practice_hours')->default(0);
            $table->integer('theory_hours')->default(0);
            $table->integer('lab_hours')->default(0);
            $table->integer('total_hours');

            $table->jsonb('class_schedule')->nullable(); // JSONB para horarios

            // Relaciones
            $table->foreignUuid('professor_id')->constrained('professors');
            $table->foreignUuid('career_id')->constrained('careers');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('subjects');
    }
};
