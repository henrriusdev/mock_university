<?php

namespace App\Http\Controllers;

use App\Models\UserInvitation;
use Illuminate\Http\Request;
use Inertia\Inertia;

class InvitationController extends Controller
{
    /**
     * Handle an incoming invitation link.
     */
    public function __invoke(Request $request, string $token)
    {
        $invitation = UserInvitation::query()
            ->with('user')
            ->where('token', $token)
            ->whereNull('consumed_at')
            ->where(function ($query) {
                $query->whereNull('expires_at')
                    ->orWhere('expires_at', '>', now());
            })
            ->first();

        if (! $invitation) {
            return Inertia::render('Invitations/Invalid', [
                'token' => $token,
            ]);
        }

        $request->session()->put('invitation_token', $token);

        return Inertia::render('Invitations/Accept', [
            'token' => $token,
            'expiresAt' => optional($invitation->expires_at)->toIso8601String(),
            'user' => [
                'email' => $invitation->user->email,
                'firstName' => $invitation->user->first_name,
                'lastName' => $invitation->user->last_name,
            ],
        ]);
    }
}
