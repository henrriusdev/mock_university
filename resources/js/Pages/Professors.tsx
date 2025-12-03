import { Pencil } from 'lucide-react';
import { useMemo, useState } from 'react';

import {
  BatchProfessorDialog,
  type ProfessorRow,
} from '@/Components/forms/BatchProfessorDialog';
import { Badge } from '@/Components/ui/badge';
import { Button } from '@/Components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/Components/ui/card';
import {
  DataTable,
  DataTableColumnHeader,
  EMPTY_FILTER_VALUE,
  type DataTableFilterConfig,
} from '@/Components/ui/data-table';
import AuthenticatedLayout from '@/Layouts/AuthenticatedLayout';
import { Head, usePage } from '@inertiajs/react';
import type { ColumnDef, Row } from '@tanstack/react-table';

const columns: ColumnDef<ProfessorRow>[] = [
  {
    accessorKey: 'firstName',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Name" />
    ),
    cell: ({ row }) => {
      const { firstName, lastName } = row.original;
      return (
        <span className="font-medium text-foreground">
          {firstName} {lastName}
        </span>
      );
    },
    meta: { title: 'Name' },
  },
  {
    accessorKey: 'email',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Email" />
    ),
    cell: ({ row }) => (
      <span className="font-mono text-xs text-muted-foreground sm:text-sm">
        {row.original.email}
      </span>
    ),
    meta: { title: 'Email' },
  },
  {
    id: 'boss',
    accessorFn: (row) => row.boss?.name ?? EMPTY_FILTER_VALUE,
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Boss" />
    ),
    cell: ({ row }) => {
      const boss = row.original.boss;
      if (!boss) {
        return <Badge variant="outline">Unassigned</Badge>;
      }

      return <span>{boss.name}</span>;
    },
    meta: { title: 'Boss' },
  },
  {
    accessorKey: 'createdAt',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Created" />
    ),
    cell: ({ row }) => {
      const created = row.original.createdAt;
      if (!created) {
        return <span className="text-muted-foreground">—</span>;
      }

      const formatted = new Date(created).toLocaleString();

      return <span className="text-muted-foreground">{formatted}</span>;
    },
    meta: { title: 'Created' },
  },
];

export default function Professors() {
  const { professors } = usePage<{ professors: ProfessorRow[] }>().props;
  const [dialogOpen, setDialogOpen] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const [selectedProfessor, setSelectedProfessor] = useState<ProfessorRow | null>(null);

  const bossFilterOptions = useMemo(() => {
    if (!professors?.length) {
      return [];
    }

    const names = new Set<string>();
    let includeUnassigned = false;

    professors.forEach((professor) => {
      if (professor.boss?.name) {
        names.add(professor.boss.name);
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
  }, [professors]);

  const filters = useMemo<DataTableFilterConfig<ProfessorRow>[]>(() => {
    if (!bossFilterOptions.length) {
      return [];
    }

    return [
      {
        columnId: 'boss',
        title: 'Boss',
        options: bossFilterOptions,
        multi: true,
      },
    ];
  }, [bossFilterOptions]);

  return (
    <AuthenticatedLayout
      header={
        <h2 className="text-xl font-semibold leading-tight text-foreground">
          Professors
        </h2>
      }
    >
      <Head title="Professors" />
      <div className="py-12">
        <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
          <Card>
            <CardHeader className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
              <CardTitle>Professors</CardTitle>
              <BatchProfessorDialog
                professor={selectedProfessor}
                isEdit={isEdit}
                open={dialogOpen}
                onOpenChange={(open) => {
                  setDialogOpen(open);
                  if (!open) {
                    setSelectedProfessor(null);
                    setIsEdit(false);
                  }
                }}
              />
            </CardHeader>
            <CardContent className="p-6">
              <DataTable
                columns={columns}
                data={professors ?? []}
                search={{
                  columnId: 'firstName',
                  placeholder: 'Search professors...',
                }}
                filters={filters}
                pageSizes={[10, 25, 50]}
                emptyMessage="No professors found."
              >
                {(row: Row<ProfessorRow>) => (
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => {
                      setSelectedProfessor(row.original);
                      setIsEdit(true);
                      setDialogOpen(true);
                    }}
                    className="gap-1"
                  >
                    <Pencil className="w-4 h-4" />
                    Edit
                  </Button>
                )}
              </DataTable>
            </CardContent>
          </Card>
        </div>
      </div>
    </AuthenticatedLayout>
  );
}
