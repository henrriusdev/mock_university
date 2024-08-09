<script>

  /** @typedef {{id: number, identityCard: string, phone: string, totalAverage: number, birthDate: string, address: string, district: string, city: string, postalCode: string, creditUnitsAccumulated: number}} Student */

  /** @typedef{{id: number, name: string, email: string, username: string, avatar: string, active: boolean}} User */
  import {
    Camera,
    ChevronLeft
  } from "lucide-svelte";

  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";

  /** @type {Student | null} */
  export let student = null;

  /** @type {User | null} */
  export let user = null;
</script>

<DirectiveLayout title={student?.id ? `Update ${user?.name}` : "Add Student"}
description="Add or update a student of the MockUniversity" keywords="add, update, inactive, mocku">
<div class="bg-muted/40 flex min-h-screen w-full flex-col">
  <div class="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
    <main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <div class="mx-auto grid max-w-[59rem] flex-1 auto-rows-max gap-4">
        <div class="flex items-center gap-4">
          <Button variant="outline" size="icon" class="h-7 w-7">
            <ChevronLeft class="h-4 w-4" />
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
        <form class="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
          <div class="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
            <Card.Root>
              <Card.Header>
                <Card.Title>Contact information</Card.Title>
                <Card.Description>
                  Fill the form with the student's contact information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 md:grid-cols-2">
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                      <Label for="phone">Phone</Label>
                      <Input id="phone" type="tel" />
                  </div>
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                      <Label for="address">Address</Label>
                      <Input id="address" type="text" />
                      <Label for="district">District</Label>
                      <Input id="district" type="text" />
                      <Label for="city">City</Label>
                      <Input id="city" type="text" />
                      <Label for="postalCode">Postal Code</Label>
                      <Input id="postalCode" type="text" />
                  </div>
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
              <Label for="identityCard">Identity Card</Label>
                      <Input id="identityCard" type="text" />
                      <Label for="birthDate">Birth Date</Label>
                      <Input id="birthDate" type="date" />
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
                <div class="grid gap-4 md:grid-cols-2">
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                    <Label for="creditUnitsAccumulated">Credit Units Accumulated</Label>
                    <Input id="creditUnitsAccumulated" type="number" />
                  </div>
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                    <Label for="totalAverage">Total Average</Label>
                    <Input id="totalAverage" type="number" />
                  </div>
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
                 <div class="grid gap-4 md:grid-cols-2">
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                    <img src={user?.avatar} alt="Avatar" class="rounded-full w-24 h-24" />
                    <Button variant="default" class="flex justify-center gap-x-3 items-center rounded-full" size="icon">
                      <Camera class="w-8 h-8 p-1" />
                    </Button>
                  </div>
                <div class="grid gap-4 md:grid-cols-2">
                  <div class="grid gap-4 md:col-span-2 md:grid-cols-3">
                    <Label for="name">Name</Label>
                    <Input id="name" type="text" />
                    <Label for="email">Email</Label>
                    <Input id="email" type="email" />
                    <Label for="username">Username</Label>
                    <Input id="username" type="text" />
                  </div>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root></Card.Root>
            <Card.Root></Card.Root>
            <Card.Root></Card.Root>
          </div>
        </form>
      </div>
    </main>
  </div>
</div>
</DirectiveLayout>
