import { useMemo } from 'react';

import { Badge } from '@/Components/ui/badge';
import { Card, CardContent } from '@/Components/ui/card';
import {
    DataTable,
    DataTableColumnHeader,
    EMPTY_FILTER_VALUE,
    type DataTableFilterConfig,
} from '@/Components/ui/data-table';
import AuthenticatedLayout from '@/Layouts/AuthenticatedLayout';
import { Head, usePage } from '@inertiajs/react';
import type { ColumnDef } from '@tanstack/react-table';
import { Button } from '@/Components/ui/button';

type CareerRow = {
    id: string;
    code: string;
    name: string;
    description: string | null;
    leader: string | null;
    leaderId: string | null;
};

const columns: ColumnDef<CareerRow>[] = [
    {
        accessorKey: 'code',
        header: ({ column }) => (
            <DataTableColumnHeader column={column} title="Code" />
        ),
        cell: ({ row }) => (
            <span className="font-mono text-xs tracking-wide uppercase text-muted-foreground sm:text-sm">
                {row.getValue<string>('code')}
            </span>
        ),
        meta: { title: 'Code' },
    },
    {
        accessorKey: 'name',
        header: ({ column }) => (
            <DataTableColumnHeader column={column} title="Name" />
        ),
        cell: ({ row }) => (
            <span className="font-medium text-foreground">
                {row.getValue<string>('name')}
            </span>
        ),
        meta: { title: 'Name' },
    },
    {
        accessorKey: 'leader',
        header: ({ column }) => (
            <DataTableColumnHeader column={column} title="Leader" />
        ),
        cell: ({ row }) => {
            const leader = row.original.leader;
            if (!leader) {
                return <Badge variant="outline">Unassigned</Badge>;
            }

            return <span>{leader}</span>;
        },
        meta: { title: 'Leader' },
    },
    {
        accessorKey: 'description',
        header: ({ column }) => (
            <DataTableColumnHeader column={column} title="Description" />
        ),
        cell: ({ row }) => {
            const description = row.original.description;
            if (!description) {
                return <span className="text-muted-foreground">—</span>;
            }

            return (
                <span className="block max-w-2xl truncate text-muted-foreground">
                    {description}
                </span>
            );
        },
        enableSorting: false,
        meta: { title: 'Description' },
    },
];

export default function Careers() {
    const { careers } = usePage<{ careers: CareerRow[] }>().props;

    const leaderFilterOptions = useMemo(() => {
        if (!careers?.length) {
            return [];
        }

        const names = new Set<string>();
        let includeUnassigned = false;

        careers.forEach((career) => {
            if (career.leader) {
                names.add(career.leader);
            } else {
                includeUnassigned = true;
            }
        });

        const options = Array.from(names)
            .sort((a, b) => a.localeCompare(b))
            .map((name) => ({ label: name, value: name }));

        if (includeUnassigned) {
            options.unshift({ label: 'Unassigned', value: EMPTY_FILTER_VALUE });
        }

        return options;
    }, [careers]);

    const filters = useMemo<DataTableFilterConfig<CareerRow>[]>(() => {
        if (!leaderFilterOptions.length) {
            return [];
        }

        return [
            {
                columnId: 'leader',
                title: 'Leader',
                options: leaderFilterOptions,
                multi: true,
            },
        ];
    }, [leaderFilterOptions]);

    return (
        <AuthenticatedLayout
            header={
                <h2 className="text-xl font-semibold leading-tight text-foreground">
                    Careers
                </h2>
            }
        >
            <Head title="Careers" />
            <div className="py-12">
                <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
                    <Card>
              <CardContent className="p-6">
                <div className="flex items-center justify-between mb-4 text-muted-foreground">
                  <h1 className="text-2xl font-semibold text-foreground">Careers</h1>
                  <Button>Add career</Button>
                </div>
                            <DataTable
                                columns={columns}
                                data={careers ?? []}
                                search={{
                                    columnId: 'name',
                                    placeholder: 'Search careers...',
                                }}
                                filters={filters}
                                pageSizes={[10, 25, 50]}
                                emptyMessage="No careers found."
                            />
                        </CardContent>
                    </Card>
                </div>
            </div>
        </AuthenticatedLayout>
    );
}
