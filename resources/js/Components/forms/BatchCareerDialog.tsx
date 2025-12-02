import axios from "axios";
import { Loader2, Plus, Trash2 } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { FieldPath, useFieldArray, useForm } from "react-hook-form";
import { router } from "@inertiajs/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

import { Button } from "../ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { Textarea } from "../ui/textarea";

type CareerRow = {
  id: string;
  code: string;
  name: string;
  description: string | null;
  leader: string | null;
  leaderId: string | null;
};

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

const createEmptyCareer = (): CareersFormItem => ({
  id: undefined,
  name: "",
  code: "",
  description: "",
  leaderId: undefined,
});

type BatchCareersDialogProps = {
  career?: CareerRow | null;
  triggerLabel?: string;
  isEdit?: boolean;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
};

export function BatchCareersDialog({
  career,
  triggerLabel = "Manage careers",
  isEdit = false,
  open: propOpen,
  onOpenChange,
}: BatchCareersDialogProps) {
  const [open, setOpen] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const initialCareers = useMemo<CareersFormItem[]>(() => {
    if (career) {
      return [
        {
          id: career.id,
          name: career.name,
          code: career.code,
          description: career.description ?? "",
          leaderId: career.leaderId ?? undefined,
        },
      ];
    }

    return [createEmptyCareer()];
  }, [career]);

  const form = useForm<CareersFormValues>({
    resolver: zodResolver(careersFormSchema),
    defaultValues: {
      careers: initialCareers,
    },
  });

  const {
    control,
    handleSubmit,
    reset,
    formState: { isSubmitting },
    setError,
  } = form;

  const { fields, append, remove } = useFieldArray({
    control,
    name: 'careers',
  });

  useEffect(() => {
    if (open) {
      reset({ careers: initialCareers });
      setApiError(null);
    }
  }, [initialCareers, open, reset]);

  useEffect(() => {
    if (!open) {
      reset({ careers: initialCareers });
      setApiError(null);
    }
  }, [initialCareers, open, reset]);

  useEffect(() => {
    if (propOpen !== undefined) {
      setOpen(propOpen);
    }
  }, [propOpen]);

  const onSubmit = handleSubmit(async (values) => {
    setApiError(null);

    const normalizedCareers = values.careers.map((careerItem: CareersFormItem) => ({
      id: careerItem.id,
      name: careerItem.name.trim(),
      code: careerItem.code.trim(),
      description:
        careerItem.description && careerItem.description.trim().length
          ? careerItem.description.trim()
          : null,
      leader_id:
        careerItem.leaderId && careerItem.leaderId.length
          ? careerItem.leaderId
          : null,
    }));

    try {
      if (isEdit) {
        const [item] = normalizedCareers;

        if (!item || !item.id) {
          setApiError('Missing career identifier for update. Please try again.');
          return;
        }

        await axios.put(route('careers.update', { career: item.id }), item);
      } else {
        await axios.post(route('careers.store'), { careers: normalizedCareers });
      }

      router.reload({
        only: ['careers'],
      });

      setOpen(false);
      if (onOpenChange) {
        onOpenChange(false);
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 422) {
          const messages = error.response.data?.errors as Record<string, string[]> | undefined;
          if (messages) {
            Object.entries(messages).forEach(([field, fieldErrors]) => {
              const message = fieldErrors?.[0];
              if (!message) {
                return;
              }

              const normalizedField = field.replace(
                'leader_id',
                'leaderId',
              ) as FieldPath<CareersFormValues>;

              setError(normalizedField, {
                type: 'server',
                message,
              });
            });
          }

          setApiError('Validation failed. Please review the highlighted fields.');
        } else {
          setApiError(
            error.response?.data?.message ??
            'Unable to save careers. Please try again.',
          );
        }
      } else {
        setApiError('Unexpected error. Please try again.');
      }
    }
  });

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button size="sm" className="inline-flex items-center gap-2">
          <Plus className="w-4 h-4" />
          {triggerLabel}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-3xl">
        <Form {...form}>
          <form onSubmit={onSubmit} className="space-y-6">
            <DialogHeader>
              <DialogTitle>Manage careers</DialogTitle>
              <DialogDescription>
                Add new careers or update existing records in one go. All
                changes are saved together.
              </DialogDescription>
            </DialogHeader>

            {apiError ? (
              <div className="px-3 py-2 text-sm border rounded-md border-destructive/40 bg-destructive/10 text-destructive">
                {apiError}
              </div>
            ) : null}

            <div className="space-y-4">
              {fields.length ? (
                fields.map((field, index) => (
                  <div
                    key={field.id}
                    className="p-4 space-y-4 border rounded-md"
                  >
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-muted-foreground">
                        {form.getValues(`careers.${index}.id`) ? 'Existing career' : 'New career'}
                      </span>
                      {fields.length > 1 ? (
                        <Button
                          type="button"
                          variant="ghost"
                          size="sm"
                          onClick={() => remove(index)}
                          disabled={isSubmitting}
                          className="text-destructive hover:text-destructive"
                        >
                          <Trash2 className="w-4 h-4 mr-2" />
                          Remove
                        </Button>
                      ) : null}
                    </div>

                    {form.getValues(`careers.${index}.id`) ? (
                      <div className="font-mono text-xs text-muted-foreground">
                        ID: {form.getValues(`careers.${index}.id`)}
                      </div>
                    ) : null}

                    <div className="grid gap-4 sm:grid-cols-2">
                      <FormField
                        control={control}
                        name={`careers.${index}.name`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Name</FormLabel>
                            <FormControl>
                              <Input
                                {...field}
                                value={field.value ?? ''}
                                placeholder="Computer Science"
                              />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={control}
                        name={`careers.${index}.code`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Code</FormLabel>
                            <FormControl>
                              <Input
                                {...field}
                                value={field.value ?? ''}
                                placeholder="CS-101"
                              />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>

                    <FormField
                      control={control}
                      name={`careers.${index}.description`}
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Description</FormLabel>
                          <FormControl>
                            <Textarea
                              {...field}
                              value={field.value ?? ''}
                              placeholder="Brief summary of the career"
                              rows={4}
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    <FormField
                      control={control}
                      name={`careers.${index}.leaderId`}
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Leader ID (optional)</FormLabel>
                          <FormControl>
                            <Input
                              {...field}
                              value={field.value ?? ''}
                              placeholder="UUID of the professor"
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </div>
                ))
              ) : (
                <div className="text-sm text-muted-foreground">
                  No careers to edit yet. Add a career to get started.
                </div>
              )}
            </div>

            <div>
              <Button
                type="button"
                variant="outline"
                onClick={() => append(createEmptyCareer())}
                disabled={isSubmitting || isEdit}
              >
                <Plus className="w-4 h-4 mr-2" />
                Add another career
              </Button>
            </div>

            <DialogFooter>
              <DialogClose asChild disabled={isSubmitting}>
                <Button type="button" variant="outline">
                  Cancel
                </Button>
              </DialogClose>
              <Button type="submit" disabled={isSubmitting}>
                {isSubmitting ? (
                  <>
                    <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                    Saving
                  </>
                ) : (
                  'Save changes'
                )}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}