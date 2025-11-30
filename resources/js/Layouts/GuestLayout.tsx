import ApplicationLogo from '@/Components/ApplicationLogo';
import { Card, CardContent } from '@/Components/ui/card';
import { Link } from '@inertiajs/react';
import { PropsWithChildren } from 'react';

export default function Guest({ children }: PropsWithChildren) {
    return (
        <div className="flex min-h-screen flex-col items-center justify-center bg-background px-4 py-10">
            <Link href="/" className="flex items-center justify-center">
                <ApplicationLogo className="h-16 w-16 fill-foreground" />
            </Link>

            <Card className="mt-8 w-full max-w-md">
                <CardContent className="p-6 sm:p-8">{children}</CardContent>
            </Card>
        </div>
    );
}
