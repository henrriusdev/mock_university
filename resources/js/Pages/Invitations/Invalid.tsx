import GuestLayout from '@/Layouts/GuestLayout';
import { Button } from '@/Components/ui/button';
import { Head, Link } from '@inertiajs/react';

interface InvalidInvitationProps {
  token: string;
}

export default function InvalidInvitation({ token }: InvalidInvitationProps) {
  return (
    <GuestLayout>
      <Head title="Invitation not found" />
      <div className="space-y-6 text-center">
        <div className="space-y-2">
          <h1 className="text-2xl font-semibold text-foreground">Invitation not available</h1>
          <p className="text-sm text-muted-foreground">
            We could not find a valid invitation for token <span className="font-mono text-xs">{token}</span>.
          </p>
        </div>

        <p className="text-sm text-muted-foreground">
          The invitation may have expired or was already used. Contact your administrator to request a new invite or
          continue to the login page if you already have access.
        </p>

        <div className="space-y-3">
          <Button asChild className="w-full">
            <Link href={route('login')}>Back to login</Link>
          </Button>
        </div>
      </div>
    </GuestLayout>
  );
}
