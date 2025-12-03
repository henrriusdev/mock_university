@component('mail::message')
# Welcome to {{ config('app.name') }}

Hello {{ $user->first_name }},

You've been invited to join {{ config('app.name') }}. Click the button below to finish setting up your account.

@component('mail::button', ['url' => $acceptUrl])
Accept Invitation
@endcomponent

@isset($expiresAt)
This link expires on {{ $expiresAt->timezone(config('app.timezone'))->toDayDateTimeString() }}.
@endisset

If you weren't expecting this invitation, you can safely ignore this message.

Thanks,
{{ config('app.name') }} Team
@endcomponent
