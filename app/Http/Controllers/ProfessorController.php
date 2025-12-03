<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreProfessorRequest;
use App\Http\Requests\UpdateProfessorRequest;
use App\Mail\UserInvitationMail;
use App\Models\Professor;
use App\Models\Role;
use App\Models\User;
use App\Models\UserInvitation;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Mail;
use Illuminate\Support\Str;
use Inertia\Inertia;

class ProfessorController extends Controller
{
    /**
     * Display a listing of the resource.
     * @group Professor
     */
    public function index()
    {
        $professors = Professor::query()
            ->with(['user', 'boss.user'])
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

        $professorRole = Role::query()
            ->where('name', 'Professor')
            ->firstOrFail();

        $createdProfessors = DB::transaction(function () use ($payload, $professorRole) {
            return $payload->map(function (array $professorData) use ($professorRole) {
                $firstName = trim($professorData['first_name']);
                $lastName = trim($professorData['last_name']);
                $fullName = trim($firstName . ' ' . $lastName);
                $email = strtolower(trim($professorData['email']));

                $user = User::query()->create([
                    'name' => $fullName,
                    'email' => $email,
                    'first_name' => $firstName,
                    'last_name' => $lastName,
                    'identification_number' => (string) Str::uuid(),
                    'password' => Hash::make(Str::random(40)),
                    'role_id' => $professorRole->id,
                ]);

                $professor = Professor::query()->create([
                    'user_id' => $user->id,
                    'boss_id' => $professorData['boss_id'] ?? null,
                ]);

                $invitation = UserInvitation::query()->create([
                    'user_id' => $user->id,
                    'token' => Str::uuid()->toString(),
                    'expires_at' => now()->addDays(7),
                ]);

                return [
                    'professor' => $professor,
                    'invitation' => $invitation,
                ];
            });
        });

        $professors = $createdProfessors->map(function (array $result) {
            /** @var Professor $professor */
            $professor = $result['professor'];
            $professor->load(['user', 'boss.user']);

            /** @var UserInvitation $invitation */
            $invitation = $result['invitation'];
            $acceptUrl = route('invitations.accept', ['token' => $invitation->token]);

            Mail::to($professor->user->email)->queue(
                new UserInvitationMail(
                    $professor->user,
                    $acceptUrl,
                    $invitation->expires_at
                )
            );

            return $professor;
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

        $firstName = trim($data['first_name']);
        $lastName = trim($data['last_name']);
        $fullName = trim($firstName . ' ' . $lastName);

        $user = $professor->user;
        $user->fill([
            'first_name' => $firstName,
            'last_name' => $lastName,
            'name' => $fullName,
            'email' => strtolower(trim($data['email'])),
        ]);
        $user->save();

        $professor->boss_id = $data['boss_id'] ?? null;
        $professor->save();

        $professor->load(['user', 'boss.user']);

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
        $user = $professor->user;

        return [
            'id' => $professor->id,
            'firstName' => $user->first_name,
            'lastName' => $user->last_name,
            'email' => $user->email,
            'identificationNumber' => $user->identification_number,
            'phone' => $user->phone,
            'address' => $user->address,
            'boss' => $professor->boss ? [
                'id' => $professor->boss->id,
                'name' => $professor->boss->user->name,
                'email' => $professor->boss->user->email,
            ] : null,
            'createdAt' => $professor->created_at?->toIso8601String(),
            'updatedAt' => $professor->updated_at?->toIso8601String(),
        ];
    }
}
