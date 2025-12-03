<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreProfessorRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'professors' => ['required', 'array', 'min:1'],
            'professors.*.first_name' => ['required', 'string', 'max:255'],
            'professors.*.last_name' => ['required', 'string', 'max:255'],
            'professors.*.boss_id' => ['nullable', 'uuid', 'exists:professors,id'],
            'professors.*.email' => ['required', 'string', 'email', 'max:255', 'unique:users,email'],
        ];
    }
}
