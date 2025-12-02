import axios from 'axios';
import { Loader2, Pencil, Plus, Trash2 } from 'lucide-react';
import { useEffect, useMemo, useState } from 'react';
import { useFieldArray, useForm, type FieldPath } from 'react-hook-form';
import { z } from 'zod';

import { Badge } from '@/Components/ui/badge';
import { Button } from '@/Components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/Components/ui/card';
import {
  DataTable,
  DataTableColumnHeader,
  EMPTY_FILTER_VALUE,
  type DataTableFilterConfig,
} from '@/Components/ui/data-table';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/Components/ui/dialog';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/Components/ui/form';
import { Input } from '@/Components/ui/input';
import { Textarea } from '@/Components/ui/textarea';
import AuthenticatedLayout from '@/Layouts/AuthenticatedLayout';
import { Head, router, usePage } from '@inertiajs/react';
import { zodResolver } from '@hookform/resolvers/zod';
import type { ColumnDef, Row } from '@tanstack/react-table';
import { BatchCareersDialog } from '@/Components/forms/BatchCareerDialog';

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

const careerFormSchema = z.object({
  id: z.string().uuid().optional(),
  name: z
    .string()
    .trim()
    .min(1, 'Name is required')
    .max(255, 'Name must not exceed 255 characters'),
  code: z
    .string()
    .trim()
    .min(1, 'Code is required')
    .max(255, 'Code must not exceed 255 characters'),
  description: z
    .string()
    .max(65535, 'Description is too long')
    .optional()
    .nullable(),
  leaderId: z
    .union([z.string().trim().length(0), z.string().uuid('Leader must be a valid UUID')])
    .optional()
    .transform((value) => (!value || value.length === 0 ? undefined : value)),
});

const careersFormSchema = z.object({
  careers: z
    .array(careerFormSchema)
    .min(1, 'Add at least one career to continue'),
});

type CareersFormValues = z.infer<typeof careersFormSchema>;
type CareersFormItem = CareersFormValues['careers'][number];

export default function Careers() {
  const { careers } = usePage<{ careers: CareerRow[] }>().props;
  const [isEdit, setIsEdit] = useState(false);
  const [career, setCareer] = useState<CareerRow | null>(null);

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
                        <CardHeader className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                            <CardTitle>Careers</CardTitle>
                            <BatchCareersDialog open={isEdit} onOpenChange={setIsEdit} isEdit={isEdit} career={career} />
                        </CardHeader>
            <CardContent className="p-6">
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
                >
                {(row: Row<CareerRow>) => (
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => {
                      setCareer(row.original);
                      setIsEdit(true);
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
