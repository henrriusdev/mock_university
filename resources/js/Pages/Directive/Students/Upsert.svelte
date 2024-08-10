<script>
  /** @typedef {{id: number, identityCard: string, phone: string, totalAverage: number, birthDate: string, address: string, district: string, city: string, postalCode: string, creditUnitsAccumulated: number, semester: number}} Student */

  /** @typedef{{id: number, name: string, email: string, username: string, avatar: string, active: boolean}} User */
  import { Camera, ChevronLeft } from "lucide-svelte";

  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input, MaskInput } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import DatePicker from "$lib/components/DatePicker.svelte";
  import Textarea from "$lib/components/ui/textarea/textarea.svelte";

  /** @type {Student | null} */
  export let student = null;

  /** @type {User | null} */
  export let user = null;
</script>

<DirectiveLayout
  title="{student?.id ? `Update ${user?.name}` : 'Add Student'}"
  description="Add or update a student of the MockUniversity"
  keywords="add, update, inactive, mocku"
>
  <section class="bg-muted/40 flex w-full flex-col p-4">
    <main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <div class="mx-auto grid max-w-[59rem] flex-1 auto-rows-max gap-4">
        <div class="flex items-center gap-4 py-1">
          <Button
            variant="ghost"
            size="icon"
            class="h-7 w-7"
            href="/directive/students"
          >
            <ChevronLeft class="h-5 w-5" />
            <span class="sr-only">Back</span>
          </Button>
          <h1
            class="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0"
          >
            {#if student?.id}
              Update {user?.name}
            {:else}
              Add Student
            {/if}
          </h1>
        </div>
        <form
          class="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8"
        >
          <div
            class="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8"
          >
            <Card.Root>
              <Card.Header>
                <Card.Title>Contact information</Card.Title>
                <Card.Description>
                  Fill the form with the student's contact information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-4">
                  <span
                    class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="phone">Phone</Label>
                    <MaskInput id="phone" type="tel" mask="(000) 000-00-00" value={student?.phone} />
                  </span>
                  <span
                  class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                  <Label for="district">District</Label>
                  <Input id="district" type="text" value={student?.district} />
                </span>
                <span
                class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                >
                <Label for="city">City</Label>
                <Input id="city" type="text" value={student?.city} />
              </span>
              <span
              class="text-sm font-semibold text-muted-foreground lg:col-span-2"
              >
              <Label for="postalCode">Postal Code</Label>
              <MaskInput id="postalCode" type="text" mask="00000" value={student?.postalCode} />
            </span>
            <span
              class="text-sm font-semibold text-muted-foreground lg:col-span-4"
            >
              <Label for="address">Address</Label>
              <Textarea id="address" value={student?.address} />
            </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Personal information</Card.Title>
                <Card.Description>
                  Fill the form with the student's personal information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-4">
                  <span
                    class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="identityCard">Identity Card</Label>
                    <MaskInput
                      id="identityCard"
                      type="text"
                      mask="V-000000000"
                      value={student?.identityCard}
                    />
                  </span>
                  <span
                    class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="birthDate">Birth Date</Label>
                    <DatePicker value={student?.birthDate} />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Academic information</Card.Title>
                <Card.Description>
                  Fill the form with the student's academic information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-4">
                  <span
                    class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="creditUnitsAccumulated"
                      >Credit Units Accumulated</Label
                    >
                    <Input id="creditUnitsAccumulated" type="number" value={student?.creditUnitsAccumulated}/>
                  </span>
                  <span
                    class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="totalAverage">Semester</Label>
                    <Input id="semester" type="number" min="1" max="10" value={student?.semester} />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
          </div>
          <div class="grid auto-rows-max items-start gap-4 lg:gap-8">
            <Card.Root>
              <Card.Header>
                <Card.Title>Student details</Card.Title>
                <Card.Description>
                  Fill the form with the student's details.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <!-- Make a image rounded of the avatar with a button for change it -->
                <div class="flex flex-col gap-4">
                  <div class="flex flex-col gap-4 items-start justify-start">
                    <img
                      src="{user?.avatar ??
                        'https://avatars.dicebear.com/api/avatars/1.svg'}"
                      alt="Avatar"
                      class="rounded-full w-24 h-24"
                    />
                    <Button
                      variant="default"
                      class="flex justify-center gap-x-3 items-center rounded-full"
                      size="icon"
                    >
                      <Camera class="w-8 h-8 p-1" />
                    </Button>
                  </div>
                  <div class="grid gap-4 w-full">
                    <span class="text-sm font-semibold text-muted-foreground">
                      <Label for="name">Name</Label>
                      <Input id="name" type="text" value={user?.name} />
                    </span>
                    <span class="text-sm font-semibold text-muted-foreground">
                      <Label for="email">Email</Label>
                      <Input id="email" type="email" value={user?.email} />
                    </span>
                    <span class="text-sm font-semibold text-muted-foreground">
                      <Label for="username">Username</Label>
                      <Input id="username" type="text" value={user?.username} />
                    </span>
                  </div>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Actions</Card.Title>
                <Card.Description>
                  Save or cancel the changes made to the student. Also inactive
                  the student.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="flex flex-col gap-4">
                  <Button variant="ghost" type="reset">Cancel</Button>
                  <Button variant="default" type="submit">
                    {student?.id ? "Update" : "Add"} student
                  </Button>
                  {#if student?.id}
                    <Button variant="destructive" type="button">
                      Inactive student
                    </Button>
                  {/if}
                </div>
              </Card.Content>
            </Card.Root>
          </div>
        </form>
      </div>
    </main>
  </section>
</DirectiveLayout>
