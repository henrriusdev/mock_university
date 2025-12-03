import GuestLayout from '@/Layouts/GuestLayout';
import { Button } from '@/Components/ui/button';
import { Head, Link } from '@inertiajs/react';

type AcceptInvitationProps = {
  token: string;
  expiresAt?: string | null;
  user: {
    email: string;
    firstName: string;
    lastName: string;
  };
};

export default function AcceptInvitation({ token, expiresAt, user }: AcceptInvitationProps) {
  const expiresLabel = expiresAt ? new Date(expiresAt).toLocaleString() : null;

  return (
    <GuestLayout>
      <Head title="Accept Invitation" />
      <div className="space-y-6">
        <div className="space-y-2 text-center">
          <h1 className="text-2xl font-semibold text-foreground">You're almost there, {user.firstName}!</h1>
          <p className="text-sm text-muted-foreground">
            We've confirmed your invitation and reserved your account using {user.email}.
          </p>
          {expiresLabel ? (
            <p className="text-xs text-muted-foreground">
              This link expires on {expiresLabel}.
            </p>
          ) : null}
        </div>

        <p className="text-sm text-muted-foreground">
          Click continue to finish creating your password and activate your access. If you already completed this
          process, you can safely close this page.
        </p>

        <Button asChild className="w-full">
          <Link href={route('register', { token })}>Continue registration</Link>
        </Button>
      </div>
    </GuestLayout>
  );
}
