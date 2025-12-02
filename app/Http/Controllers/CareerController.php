<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreCareerRequest;
use App\Http\Requests\UpdateCareerRequest;
use App\Models\Career;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;
use Inertia\Inertia;

class CareerController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $careers = Career::query()
            ->with(['leader.user'])
            ->get()
            ->map(fn (Career $career) => $this->transformCareer($career))
            ->values();

        return Inertia::render('Careers', [
            'careers' => $careers,
        ]);
    }

    /**
     * Store batch careers.
     */
    public function store(StoreCareerRequest $request)
    {
        $payload = collect($request->validated('careers'));

        $careers = DB::transaction(function () use ($payload) {
            return $payload->map(function (array $careerData) {
                $description = $careerData['description'] ?? null;
                if (is_string($description)) {
                    $description = trim($description);
                    $description = $description === '' ? null : $description;
                }

                $leaderId = $careerData['leader_id'] ?? null;
                if (is_string($leaderId) && $leaderId === '') {
                    $leaderId = null;
                }

                $attributes = [
                    'name' => trim($careerData['name']),
                    'code' => trim($careerData['code']),
                    'description' => $description,
                    'leader_id' => $leaderId,
                ];

                if (! empty($careerData['id'])) {
                    $career = Career::query()->findOrFail($careerData['id']);
                    $career->fill($attributes);
                    $career->save();

                    return $career->load('leader.user');
                }

                $career = Career::query()->create([
                    'id' => (string) Str::uuid(),
                    ...$attributes,
                ]);

                return $career->load('leader.user');
            });
        });

        $response = $careers
            ->map(fn (Career $career) => $this->transformCareer($career))
            ->values();

        return response()->json([
            'careers' => $response,
        ]);
    }

    /**
     * Display the specified career.
     */
    public function show(Career $career)
    {
        //
    }

    /**
     * Update the specified career in storage.
     */
    public function update(UpdateCareerRequest $request, Career $career)
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

        $career->fill([
            'name' => trim($data['name']),
            'code' => trim($data['code']),
            'description' => $description,
            'leader_id' => $leaderId,
        ]);

        $career->save();

        $career->load('leader.user');

        return response()->json([
            'career' => $this->transformCareer($career),
        ]);
    }

    /**
     * Remove the specified career from storage.
     */
    public function destroy(Career $career)
    {
        //
    }

    /**
     * @return array<string, mixed>
     */
    private function transformCareer(Career $career): array
    {
        $career->loadMissing(['leader.user']);

        $leaderName = null;

        if ($career->leader?->user) {
            $user = $career->leader->user;

            if (! empty($user->name)) {
                $leaderName = $user->name;
            } else {
                $parts = array_values(array_filter([
                    $user->first_name,
                    $user->last_name,
                ]));

                $leaderName = $parts ? implode(' ', $parts) : null;
            }
        }

        return [
            'id' => $career->id,
            'code' => $career->code,
            'name' => $career->name,
            'description' => $career->description,
            'leader' => $leaderName,
            'leaderId' => $career->leader_id,
        ];
    }
}
