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
        Schema::create('configurations', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->dateTime('start_registration_subjects');
            $table->dateTime('end_registration_subjects');
            $table->boolean('block_not_pay_inscription')->default(false);
            $table->jsonb('fee_dates')->nullable();
            $table->integer('number_fees');
            $table->integer('number_notes');
            $table->jsonb('notes_percentages')->nullable();

            $table->foreignUuid('cycle_id')->constrained('cycles');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('configurations');
    }
};
