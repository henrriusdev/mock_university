<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreCareerRequest;
use App\Http\Requests\UpdateCareerRequest;
use App\Models\Career;
use Inertia\Inertia;

class CareerController extends Controller
{
    /**
     * Display a listing of the resource.
     * @group Career
     */
    public function index()
    {
        $careers = Career::query()
            ->with(['leader.user'])
            ->get()
            ->map(function (Career $career) {
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
            })
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
        //
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
        //
    }

    /**
     * Remove the specified career from storage.
     */
    public function destroy(Career $career)
    {
        //
    }
}
