<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreProfessorRequest;
use App\Http\Requests\UpdateProfessorRequest;
use App\Models\Professor;
use Illuminate\Support\Facades\DB;
use Inertia\Inertia;
use Str;

class ProfessorController extends Controller
{
    /**
     * Display a listing of the resource.
     * @group Professor
     */
    public function index()
    {
        $professors = Professor::query()
            ->with(['leader.user'])
            ->get()
            ->map(fn(Professor $professor) => $this->transformProfessor($professor))
            ->values();

        return Inertia::render('Professors', [
            'professors' => $professors,
        ]);
    }

    /**
     * Store batch professors.
     */
    public function store(StoreProfessorRequest $request)
    {
        $payload = collect($request->validated('professors'));

        $professors = DB::transaction(function () use ($payload) {
            return $payload->map(function (array $professorData) {
                $description = $professorData['description'] ?? null;
                if (is_string($description)) {
                    $description = trim($description);
                    $description = $description === '' ? null : $description;
                }

                $leaderId = $professorData['leader_id'] ?? null;
                if (is_string($leaderId) && $leaderId === '') {
                    $leaderId = null;
                }

                $attributes = [
                    'name' => trim($professorData['name']),
                    'code' => trim($professorData['code']),
                    'description' => $description,
                    'leader_id' => $leaderId,
                ];

                if (! empty($professorData['id'])) {
                    $professor = Professor::query()->findOrFail($professorData['id']);
                    $professor->fill($attributes);
                    $professor->save();

                    return $professor->load('leader.user');
                }

                $professor = Professor::query()->create([
                    'id' => (string) Str::uuid(),
                    ...$attributes,
                ]);

                return $professor->load('leader.user');
            });
        });

        $response = $professors
            ->map(fn(Professor $professor) => $this->transformProfessor($professor))
            ->values();

        return response()->json([
            'professors' => $response,
        ]);
    }

    /**
     * Display the specified professor.
     */
    public function show(Professor $professor)
    {
        //
    }

    /**
     * Update the specified professor in storage.
     */
    public function update(UpdateProfessorRequest $request, Professor $professor)
    {
        $data = $request->validated();

        $description = $data['description'] ?? null;
        if (is_string($description)) {
            $description = trim($description);
            $description = $description === '' ? null : $description;
        }

        $leaderId = $data['leader_id'] ?? null;
        if (is_string($leaderId) && $leaderId === '') {
            $leaderId = null;
        }

        $professor->fill([
            'name' => trim($data['name']),
            'code' => trim($data['code']),
            'description' => $description,
            'leader_id' => $leaderId,
        ]);

        $professor->save();

        $professor->load('leader.user');

        return response()->json([
            'professor' => $this->transformProfessor($professor),
        ]);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Professor $professor)
    {
        //
    }

    public function transformProfessor(Professor $professor): array
    {
        return [
            'id' => $professor->id,
            'name' => $professor->name,
            'code' => $professor->code,
            'description' => $professor->description,
            'leader' => $professor->leader ? [
                'id' => $professor->leader->id,
                'user' => [
                    'id' => $professor->leader->user->id,
                    'name' => $professor->leader->user->name,
                    'email' => $professor->leader->user->email,
                ],
            ] : null,
            'created_at' => $professor->created_at->toDateTimeString(),
            'updated_at' => $professor->updated_at->toDateTimeString(),
        ];
    }
}
