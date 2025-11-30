import {
    type Column,
    type ColumnDef,
    type ColumnFiltersState,
    type FilterFn,
    type SortingState,
    type Table as TanTable,
    type VisibilityState,
    flexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table';
import {
    ArrowDown,
    ArrowUp,
    ArrowUpDown,
    ChevronDown,
    Filter,
    X,
} from 'lucide-react';
import * as React from 'react';

import { cn } from '@/lib/utils';

import { Badge } from '@/Components/ui/badge';
import { Button } from '@/Components/ui/button';
import {
    DropdownMenu,
    DropdownMenuCheckboxItem,
    DropdownMenuContent,
    DropdownMenuRadioGroup,
    DropdownMenuRadioItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/Components/ui/dropdown-menu';
import { Input } from '@/Components/ui/input';
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/Components/ui/select';
import {
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
    Table as UiTable,
} from '@/Components/ui/table';

export const EMPTY_FILTER_VALUE = '__EMPTY__';

export type DataTableFilterOption = {
    label: string;
    value: string;
    icon?: React.ComponentType<{ className?: string }>;
};

export type DataTableFilterConfig<TData> = {
    columnId: string;
    title: string;
    options: DataTableFilterOption[];
    multi?: boolean;
    filterFn?: FilterFn<TData>;
};

type DataTableSearchConfig = {
    columnId: string;
    placeholder?: string;
};

interface DataTableProps<TData, TValue> {
    columns: ColumnDef<TData, TValue>[];
    data: TData[];
    search?: DataTableSearchConfig;
    filters?: Array<DataTableFilterConfig<TData>>;
    pageSizes?: number[];
    defaultPageSize?: number;
    defaultColumnVisibility?: VisibilityState;
    emptyMessage?: string;
    className?: string;
}

export function DataTable<TData, TValue>({
    columns,
    data,
    search,
    filters = [],
    pageSizes = [10, 20, 50],
    defaultPageSize,
    defaultColumnVisibility,
    emptyMessage = 'No results found.',
    className,
}: DataTableProps<TData, TValue>) {
    const [sorting, setSorting] = React.useState<SortingState>([]);
    const [columnFilters, setColumnFilters] =
        React.useState<ColumnFiltersState>([]);
    const [columnVisibility, setColumnVisibility] =
        React.useState<VisibilityState>(defaultColumnVisibility ?? {});
    const [rowSelection, setRowSelection] = React.useState({});

    const initialPageSize = React.useMemo(() => {
        const sizes = pageSizes.length > 0 ? pageSizes : [10];
        const candidate = defaultPageSize ?? sizes[0];
        return candidate > 0 ? candidate : 10;
    }, [pageSizes, defaultPageSize]);

    const filterConfigById = React.useMemo(() => {
        return new Map(filters.map((filter) => [filter.columnId, filter]));
    }, [filters]);

    const multiSelectFilterFn = React.useCallback<FilterFn<TData>>(
        (row, columnId, filterValue) => {
            const selections = Array.isArray(filterValue)
                ? (filterValue as unknown[]).map((value) => String(value))
                : [];

            if (selections.length === 0) {
                return true;
            }

            const rawValue = row.getValue(columnId);
            const values = Array.isArray(rawValue)
                ? rawValue.map((value) =>
                      value == null || value === ''
                          ? EMPTY_FILTER_VALUE
                          : String(value),
                  )
                : [
                      rawValue == null || rawValue === ''
                          ? EMPTY_FILTER_VALUE
                          : String(rawValue),
                  ];

            return selections.some((selection) => values.includes(selection));
        },
        [],
    );

    const equalsFilterFn = React.useCallback<FilterFn<TData>>(
        (row, columnId, filterValue) => {
            const nextValue = Array.isArray(filterValue)
                ? (filterValue as unknown[])[0]
                : filterValue;

            if (nextValue == null || nextValue === '') {
                return true;
            }

            const rawValue = row.getValue(columnId);
            const comparable =
                rawValue == null || rawValue === ''
                    ? EMPTY_FILTER_VALUE
                    : String(rawValue);
            return comparable === String(nextValue);
        },
        [],
    );

    const resolvedColumns = React.useMemo(() => {
        if (!filterConfigById.size) {
            return columns;
        }

        const patchColumns = <TVal,>(
            defs: ColumnDef<TData, TVal>[],
        ): ColumnDef<TData, TVal>[] =>
            defs.map((column) => {
                if ('columns' in column && column.columns) {
                    return {
                        ...column,
                        columns: patchColumns(
                            column.columns as ColumnDef<TData, unknown>[],
                        ),
                    };
                }

                const accessorKey =
                    'accessorKey' in column &&
                    typeof column.accessorKey === 'string'
                        ? column.accessorKey
                        : undefined;
                const columnId = column.id ?? accessorKey;
                if (!columnId) {
                    return column;
                }

                const config = filterConfigById.get(columnId);
                if (!config) {
                    return column;
                }

                if (column.filterFn) {
                    return column;
                }

                if (config.filterFn) {
                    return {
                        ...column,
                        filterFn: config.filterFn,
                    };
                }

                if (config.multi) {
                    return {
                        ...column,
                        filterFn: multiSelectFilterFn,
                    };
                }

                return {
                    ...column,
                    filterFn: equalsFilterFn,
                };
            });

        return patchColumns(
            columns as ColumnDef<TData, unknown>[],
        ) as ColumnDef<TData, TValue>[];
    }, [columns, equalsFilterFn, filterConfigById, multiSelectFilterFn]);

    const table = useReactTable({
        data,
        columns: resolvedColumns,
        state: {
            sorting,
            columnFilters,
            columnVisibility,
            rowSelection,
        },
        onSortingChange: setSorting,
        onColumnFiltersChange: setColumnFilters,
        onColumnVisibilityChange: setColumnVisibility,
        onRowSelectionChange: setRowSelection,
        getCoreRowModel: getCoreRowModel(),
        getSortedRowModel: getSortedRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        initialState: {
            pagination: {
                pageSize: initialPageSize,
            },
        },
    });

    return (
        <div className={cn('space-y-4', className)}>
            <DataTableToolbar table={table} search={search} filters={filters} />
            <div className="rounded-md border">
                <UiTable>
                    <TableHeader>
                        {table.getHeaderGroups().map((headerGroup) => (
                            <TableRow key={headerGroup.id}>
                                {headerGroup.headers.map((header) => (
                                    <TableHead key={header.id}>
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(
                                                  header.column.columnDef
                                                      .header,
                                                  header.getContext(),
                                              )}
                                    </TableHead>
                                ))}
                            </TableRow>
                        ))}
                    </TableHeader>
                    <TableBody>
                        {table.getRowModel().rows.length ? (
                            table.getRowModel().rows.map((row) => (
                                <TableRow
                                    key={row.id}
                                    data-state={
                                        row.getIsSelected()
                                            ? 'selected'
                                            : undefined
                                    }
                                >
                                    {row.getVisibleCells().map((cell) => (
                                        <TableCell key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext(),
                                            )}
                                        </TableCell>
                                    ))}
                                </TableRow>
                            ))
                        ) : (
                            <TableRow>
                                <TableCell
                                    colSpan={table.getAllLeafColumns().length}
                                    className="h-24 text-center text-sm text-muted-foreground"
                                >
                                    {emptyMessage}
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </UiTable>
            </div>
            <DataTablePagination
                table={table}
                pageSizes={pageSizes.length ? pageSizes : [initialPageSize]}
            />
        </div>
    );
}

type DataTableToolbarProps<TData> = {
    table: TanTable<TData>;
    search?: DataTableSearchConfig;
    filters?: Array<DataTableFilterConfig<TData>>;
};

function DataTableToolbar<TData>({
    table,
    search,
    filters = [],
}: DataTableToolbarProps<TData>) {
    const searchColumn = search ? table.getColumn(search.columnId) : undefined;
    const isFiltered = table.getState().columnFilters.length > 0;

    return (
        <div className="flex flex-col gap-3 py-4 sm:flex-row sm:items-center sm:justify-between">
            <div className="flex flex-1 flex-wrap items-center gap-2">
                {searchColumn ? (
                    <Input
                        placeholder={search?.placeholder ?? 'Search...'}
                        value={(searchColumn.getFilterValue() as string) ?? ''}
                        onChange={(event) =>
                            searchColumn.setFilterValue(
                                event.target.value
                                    ? event.target.value
                                    : undefined,
                            )
                        }
                        className="h-9 w-full sm:w-64"
                    />
                ) : null}

                {filters.map((filter) => {
                    const column = table.getColumn(filter.columnId);
                    if (!column || !column.getCanFilter()) {
                        return null;
                    }

                    return (
                        <DataTableFacetedFilter
                            key={filter.columnId}
                            column={column}
                            title={filter.title}
                            options={filter.options}
                            multi={filter.multi}
                        />
                    );
                })}

                {isFiltered ? (
                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() => table.resetColumnFilters()}
                        className="h-8 px-2 text-sm"
                    >
                        Reset
                        <X className="ml-2 h-4 w-4" />
                    </Button>
                ) : null}
            </div>

            <DataTableViewOptions table={table} />
        </div>
    );
}

type DataTableFacetedFilterProps<TData> = {
    column: Column<TData, unknown>;
    title: string;
    options: DataTableFilterOption[];
    multi?: boolean;
};

function DataTableFacetedFilter<TData>({
    column,
    title,
    options,
    multi,
}: DataTableFacetedFilterProps<TData>) {
    const rawFilterValue = column.getFilterValue();
    const selectedValues = React.useMemo(() => {
        if (multi) {
            const values = Array.isArray(rawFilterValue)
                ? (rawFilterValue as string[])
                : rawFilterValue
                  ? [String(rawFilterValue)]
                  : [];
            return new Set(values);
        }

        return rawFilterValue
            ? new Set([String(rawFilterValue)])
            : new Set<string>();
    }, [multi, rawFilterValue]);

    const selectedValue = React.useMemo(
        () => (selectedValues.size ? Array.from(selectedValues)[0] : ''),
        [selectedValues],
    );

    const clearFilter = React.useCallback(() => {
        column.setFilterValue(undefined);
    }, [column]);

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    variant="outline"
                    size="sm"
                    className="h-8 border-dashed"
                >
                    <Filter className="mr-2 h-4 w-4" />
                    {title}
                    {selectedValues.size > 0 ? (
                        <Badge
                            variant="secondary"
                            className="ml-2 rounded-full px-2"
                        >
                            {selectedValues.size}
                        </Badge>
                    ) : null}
                    <ChevronDown className="ml-2 h-4 w-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="start" className="w-60 p-1">
                {multi ? (
                    options.map((option) => {
                        const checked = selectedValues.has(option.value);
                        return (
                            <DropdownMenuCheckboxItem
                                key={option.value}
                                checked={checked}
                                onCheckedChange={(next) => {
                                    const nextValues = new Set(selectedValues);
                                    if (next) {
                                        nextValues.add(option.value);
                                    } else {
                                        nextValues.delete(option.value);
                                    }
                                    const arrayValues = Array.from(nextValues);
                                    column.setFilterValue(
                                        arrayValues.length
                                            ? arrayValues
                                            : undefined,
                                    );
                                }}
                                className="capitalize"
                            >
                                {option.icon ? (
                                    <option.icon className="mr-2 h-4 w-4" />
                                ) : null}
                                {option.label}
                            </DropdownMenuCheckboxItem>
                        );
                    })
                ) : (
                    <DropdownMenuRadioGroup
                        value={selectedValue}
                        onValueChange={(value) => {
                            if (!value) {
                                clearFilter();
                                return;
                            }

                            column.setFilterValue(value);
                        }}
                        className="capitalize"
                    >
                        {options.map((option) => (
                            <DropdownMenuRadioItem
                                key={option.value}
                                value={option.value}
                                className="capitalize"
                            >
                                <span className="flex w-full items-center gap-2">
                                    {option.icon ? (
                                        <option.icon className="h-4 w-4" />
                                    ) : null}
                                    <span>{option.label}</span>
                                </span>
                            </DropdownMenuRadioItem>
                        ))}
                    </DropdownMenuRadioGroup>
                )}

                {selectedValues.size > 0 ? (
                    <>
                        <DropdownMenuSeparator />
                        <div className="flex justify-end px-1.5 py-1">
                            <Button
                                variant="ghost"
                                size="sm"
                                onClick={clearFilter}
                            >
                                Clear
                            </Button>
                        </div>
                    </>
                ) : null}
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

type DataTableViewOptionsProps<TData> = {
    table: TanTable<TData>;
};

function DataTableViewOptions<TData>({
    table,
}: DataTableViewOptionsProps<TData>) {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="outline" size="sm" className="ml-auto h-8">
                    Columns
                    <ChevronDown className="ml-2 h-4 w-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-48">
                {table
                    .getAllLeafColumns()
                    .filter((column) => column.getCanHide())
                    .map((column) => (
                        <DropdownMenuCheckboxItem
                            key={column.id}
                            className="capitalize"
                            checked={column.getIsVisible()}
                            onCheckedChange={(checked) => {
                                column.toggleVisibility(Boolean(checked));
                            }}
                        >
                            {(
                                column.columnDef.meta as
                                    | { title?: string }
                                    | undefined
                            )?.title ?? column.id}
                        </DropdownMenuCheckboxItem>
                    ))}
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

type DataTablePaginationProps<TData> = {
    table: TanTable<TData>;
    pageSizes: number[];
};

function DataTablePagination<TData>({
    table,
    pageSizes,
}: DataTablePaginationProps<TData>) {
    const { pageIndex, pageSize } = table.getState().pagination;
    const totalRows = table.getFilteredRowModel().rows.length;
    const currentRows = table.getRowModel().rows.length;
    const pageStart = totalRows === 0 ? 0 : pageIndex * pageSize + 1;
    const pageEnd = totalRows === 0 ? 0 : pageStart + currentRows - 1;

    return (
        <div className="flex flex-col gap-3 py-2 sm:flex-row sm:items-center sm:justify-between">
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <span>Rows per page</span>
                <Select
                    value={String(pageSize)}
                    onValueChange={(value) => {
                        table.setPageSize(Number(value));
                    }}
                >
                    <SelectTrigger className="h-8 w-24">
                        <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                        {pageSizes.map((size) => (
                            <SelectItem key={size} value={String(size)}>
                                {size}
                            </SelectItem>
                        ))}
                    </SelectContent>
                </Select>
                <span className="hidden sm:inline">
                    Showing {pageStart}-{pageEnd} of {totalRows}
                </span>
            </div>
            <div className="flex items-center gap-2">
                <span className="text-sm text-muted-foreground">
                    Page {table.getPageCount() === 0 ? 0 : pageIndex + 1} of{' '}
                    {table.getPageCount()}
                </span>
                <Button
                    variant="outline"
                    size="sm"
                    onClick={() => table.previousPage()}
                    disabled={!table.getCanPreviousPage()}
                >
                    Previous
                </Button>
                <Button
                    variant="outline"
                    size="sm"
                    onClick={() => table.nextPage()}
                    disabled={!table.getCanNextPage()}
                >
                    Next
                </Button>
            </div>
        </div>
    );
}

type DataTableColumnHeaderProps<TData, TValue> = {
    column: Column<TData, TValue>;
    title: React.ReactNode;
    className?: string;
};

export function DataTableColumnHeader<TData, TValue>({
    column,
    title,
    className,
}: DataTableColumnHeaderProps<TData, TValue>) {
    if (!column.getCanSort()) {
        return (
            <div className={cn('text-sm font-medium', className)}>{title}</div>
        );
    }

    const isSorted = column.getIsSorted();

    return (
        <Button
            variant="ghost"
            size="sm"
            className={cn(
                'flex h-8 items-center justify-start gap-2 px-2 text-sm font-medium',
                className,
            )}
            onClick={() => column.toggleSorting(isSorted === 'asc')}
        >
            <span>{title}</span>
            {isSorted === 'asc' ? (
                <ArrowUp className="h-4 w-4" />
            ) : isSorted === 'desc' ? (
                <ArrowDown className="h-4 w-4" />
            ) : (
                <ArrowUpDown className="h-4 w-4" />
            )}
        </Button>
    );
}
