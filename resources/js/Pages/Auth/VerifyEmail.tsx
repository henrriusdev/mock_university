import { Button } from '@/Components/ui/button';
import GuestLayout from '@/Layouts/GuestLayout';
import { Head, Link, useForm } from '@inertiajs/react';
import { FormEventHandler } from 'react';

export default function VerifyEmail({ status }: { status?: string }) {
    const { post, processing } = useForm({});

    const submit: FormEventHandler = (e) => {
        e.preventDefault();

        post(route('verification.send'));
    };

    return (
        <GuestLayout>
            <Head title="Email Verification" />

            <div className="mb-4 text-sm text-gray-600 dark:text-gray-400">
                Thanks for signing up! Before getting started, could you verify
                your email address by clicking on the link we just emailed to
                you? If you didn't receive the email, we will gladly send you
                another.
            </div>

            {status === 'verification-link-sent' && (
                <div className="mb-6 rounded-md border border-emerald-200 bg-emerald-50 px-4 py-2 text-sm text-emerald-700 dark:border-emerald-900/50 dark:bg-emerald-950/40 dark:text-emerald-200">
                    A new verification link has been sent to the email address
                    you provided during registration.
                </div>
            )}

            <form
                onSubmit={submit}
                className="mt-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
            >
                <Button type="submit" disabled={processing}>
                    Resend Verification Email
                </Button>

                <Button variant="link" asChild>
                    <Link href={route('logout')} method="post" as="button">
                        Log Out
                    </Link>
                </Button>
            </form>
        </GuestLayout>
    );
}
