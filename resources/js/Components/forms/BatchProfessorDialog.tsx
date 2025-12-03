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

export type ProfessorRow = {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
  boss: {
    id: string;
    name: string;
    email: string;
  } | null;
};

const professorSchema = z.object({
  id: z.string().uuid().optional(),
  firstName: z
    .string()
    .trim()
    .min(1, "First name is required")
    .max(255, "First name must not exceed 255 characters"),
  lastName: z
    .string()
    .trim()
    .min(1, "Last name is required")
    .max(255, "Last name must not exceed 255 characters"),
  email: z
    .string()
    .trim()
    .email("Email must be a valid address")
    .max(255, "Email must not exceed 255 characters"),
  bossId: z
    .union([z.string().trim().length(0), z.string().uuid("Boss must be a valid UUID")])
    .transform((value) => (!value || value.length === 0 ? undefined : value))
    .optional(),
});

const professorsSchema = z.object({
  professors: z.array(professorSchema).min(1, "Add at least one professor to continue"),
});

export type ProfessorsFormValues = z.infer<typeof professorsSchema>;
export type ProfessorsFormItem = ProfessorsFormValues["professors"][number];

const createEmptyProfessor = (): ProfessorsFormItem => ({
  id: undefined,
  firstName: "",
  lastName: "",
  email: "",
  bossId: undefined,
});

type BatchProfessorDialogProps = {
  professor?: ProfessorRow | null;
  triggerLabel?: string;
  isEdit?: boolean;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
};

export function BatchProfessorDialog({
  professor,
  triggerLabel = "Invite professors",
  isEdit = false,
  open: propOpen,
  onOpenChange,
}: BatchProfessorDialogProps) {
  const [open, setOpen] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const initialProfessors = useMemo<ProfessorsFormItem[]>(() => {
    if (professor) {
      return [
        {
          id: professor.id,
          firstName: professor.firstName,
          lastName: professor.lastName,
          email: professor.email,
          bossId: professor.boss?.id,
        },
      ];
    }

    return [createEmptyProfessor()];
  }, [professor]);

  const form = useForm<ProfessorsFormValues>({
    resolver: zodResolver(professorsSchema),
    defaultValues: {
      professors: initialProfessors,
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
    name: "professors",
  });

  useEffect(() => {
    if (propOpen !== undefined) {
      setOpen(propOpen);
    }
  }, [propOpen]);

  useEffect(() => {
    if (open) {
      reset({ professors: initialProfessors });
      setApiError(null);
    }
  }, [initialProfessors, open, reset]);

  useEffect(() => {
    if (!open) {
      reset({ professors: initialProfessors });
      setApiError(null);
    }
  }, [initialProfessors, open, reset]);

  const onSubmit = handleSubmit(async (values) => {
    setApiError(null);

    const normalized = values.professors.map((entry) => ({
      id: entry.id,
      first_name: entry.firstName.trim(),
      last_name: entry.lastName.trim(),
      email: entry.email.trim().toLowerCase(),
      boss_id: entry.bossId && entry.bossId.length ? entry.bossId : null,
    }));

    try {
      if (isEdit) {
        const [item] = normalized;

        if (!item || !item.id) {
          setApiError("Missing professor identifier for update. Please try again.");
          return;
        }

        await axios.put(route("professors.update", { professor: item.id }), item);
      } else {
        await axios.post(route("professors.store"), { professors: normalized });
      }

      router.reload({
        only: ["professors"],
      });

      setOpen(false);
      onOpenChange?.(false);
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

              const normalizedField = field
                .replace(/first_name$/, "firstName")
                .replace(/last_name$/, "lastName")
                .replace(/boss_id$/, "bossId") as FieldPath<ProfessorsFormValues>;

              setError(normalizedField, {
                type: "server",
                message,
              });
            });
          }

          setApiError("Validation failed. Please review the highlighted fields.");
        } else {
          setApiError(
            error.response?.data?.message ??
              "Unable to save professors. Please try again."
          );
        }
      } else {
        setApiError("Unexpected error. Please try again.");
      }
    }
  });

  return (
    <Dialog open={open} onOpenChange={(value) => {
      setOpen(value);
      onOpenChange?.(value);
    }}>
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
              <DialogTitle>{isEdit ? "Update professor" : "Invite professors"}</DialogTitle>
              <DialogDescription>
                {isEdit
                  ? "Edit the selected professor's basic information."
                  : "Add one or more professors and we will email them an invitation to finish setting up their accounts."}
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
                  <div key={field.id} className="p-4 space-y-4 border rounded-md">
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-muted-foreground">
                        {form.getValues(`professors.${index}.id`) ? "Existing professor" : "New invite"}
                      </span>
                      {!isEdit && fields.length > 1 ? (
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

                    <div className="grid gap-4 sm:grid-cols-2">
                      <FormField
                        control={control}
                        name={`professors.${index}.firstName`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>First name</FormLabel>
                            <FormControl>
                              <Input {...field} value={field.value ?? ""} placeholder="Ada" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={control}
                        name={`professors.${index}.lastName`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Last name</FormLabel>
                            <FormControl>
                              <Input {...field} value={field.value ?? ""} placeholder="Lovelace" />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>

                    <FormField
                      control={control}
                      name={`professors.${index}.email`}
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Email</FormLabel>
                          <FormControl>
                            <Input {...field} value={field.value ?? ""} placeholder="ada@example.com" type="email" />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    <FormField
                      control={control}
                      name={`professors.${index}.bossId`}
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Boss ID (optional)</FormLabel>
                          <FormControl>
                            <Input {...field} value={field.value ?? ""} placeholder="UUID of their supervisor" />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </div>
                ))
              ) : (
                <div className="text-sm text-muted-foreground">No professors to edit yet. Add a professor to get started.</div>
              )}
            </div>

            {!isEdit ? (
              <div>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => append(createEmptyProfessor())}
                  disabled={isSubmitting}
                >
                  <Plus className="w-4 h-4 mr-2" />
                  Add another professor
                </Button>
              </div>
            ) : null}

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
                  "Save changes"
                )}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
