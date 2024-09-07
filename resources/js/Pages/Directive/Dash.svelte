<script>
  import { Activity, ArrowUpRight, CreditCard, Users } from "lucide-svelte";

  import {
    Avatar,
    AvatarFallback,
    AvatarImage,
  } from "$lib/components/ui/avatar/index.js";
  import { Badge } from "$lib/components/ui/badge/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
  } from "$lib/components/ui/card/index.js";
  import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
  } from "$lib/components/ui/table/index.js";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
</script>

<DirectiveLayout
  title="Dashboard"
  description="Dashboard for Acme Inc."
  keywords="dashboard, acme, inc"
>
  <div class="grid gap-4 md:grid-cols-2 md:gap-8 lg:grid-cols-4">
    <Card>
      <CardHeader
        class="flex flex-row items-center justify-between space-y-0 pb-2"
      >
        <CardTitle class="text-sm font-medium">Total Students</CardTitle>
        <Users class="h-4 w-4 text-muted-foreground" />
      </CardHeader>
      <CardContent>
        <div class="text-2xl font-bold">48593</div>
        <p class="text-xs text-muted-foreground">
          124 already payed the inscription
        </p>
      </CardContent>
    </Card>
    <Card>
      <CardHeader
        class="flex flex-row items-center justify-between space-y-0 pb-2"
      >
        <CardTitle class="text-sm font-medium">Professors</CardTitle>
        <Users class="h-4 w-4 text-muted-foreground" />
      </CardHeader>
      <CardContent>
        <div class="text-2xl font-bold">2350</div>
        <p class="text-xs text-muted-foreground">129 are news</p>
      </CardContent>
    </Card>
    <Card>
      <CardHeader
        class="flex flex-row items-center justify-between space-y-0 pb-2"
      >
        <CardTitle class="text-sm font-medium">Payments</CardTitle>
        <CreditCard class="h-4 w-4 text-muted-foreground" />
      </CardHeader>
      <CardContent>
        <div class="text-2xl font-bold">+12,234</div>
        <p class="text-xs text-muted-foreground">+19% from last month</p>
      </CardContent>
    </Card>
    <Card>
      <CardHeader
        class="flex flex-row items-center justify-between space-y-0 pb-2"
      >
        <CardTitle class="text-sm font-medium">Requests</CardTitle>
        <Activity class="h-4 w-4 text-muted-foreground" />
      </CardHeader>
      <CardContent>
        <div class="text-2xl font-bold">+573</div>
        <p class="text-xs text-muted-foreground">+201 since last hour</p>
      </CardContent>
    </Card>
  </div>
  <div class="grid gap-4 md:gap-8 lg:grid-cols-2 xl:grid-cols-3">
    <Card class="xl:col-span-2">
      <CardHeader class="flex flex-row items-center">
        <div class="grid gap-2">
          <CardTitle>Payments</CardTitle>
          <CardDescription>Recent payments from your students.</CardDescription>
        </div>
        <Button href="/payments" size="sm" class="ml-auto gap-1">
          View All
          <ArrowUpRight class="h-4 w-4" />
        </Button>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Student</TableHead>
              <TableHead class="xl:table.-column hidden">Type</TableHead>
              <TableHead class="xl:table.-column hidden">Status</TableHead>
              <TableHead class="xl:table.-column hidden">Date</TableHead>
              <TableHead class="text-right">Amount</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>
                <div class="font-medium">Liam Johnson</div>
                <div class="hidden text-sm text-muted-foreground md:inline">
                  liam@example.com
                </div>
              </TableCell>
              <TableCell class="xl:table.-column hidden">Sale</TableCell>
              <TableCell class="xl:table.-column hidden">
                <Badge class="text-xs" variant="outline">Approved</Badge>
              </TableCell>
              <TableCell
                class="md:table.-cell xl:table.-column hidden lg:hidden"
              >
                2023-06-23
              </TableCell>
              <TableCell class="text-right">$250.00</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>
    <Card>
      <CardHeader>
        <CardTitle>Recent Requests</CardTitle>
      </CardHeader>
      <CardContent class="grid gap-8">
        <div class="flex items-center gap-4">
          <Avatar class="hidden h-9 w-9 sm:flex">
            <AvatarImage src="/avatars/01.png" alt="Avatar" />
            <AvatarFallback>OM</AvatarFallback>
          </Avatar>
          <div class="grid gap-1">
            <p class="text-sm font-medium leading-none">Olivia Martin</p>
            <p class="text-sm text-muted-foreground">olivia.martin@email.com</p>
          </div>
          <div class="ml-auto font-medium">Hello, how are you?</div>
        </div>
      </CardContent>
    </Card>
  </div>
</DirectiveLayout>
