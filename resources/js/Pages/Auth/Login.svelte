<script>
  import { Button } from "$lib/components/ui/button/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import MainLayout from "$lib/layouts/MainLayout.svelte";

  /** @type {Array<{name: string}>} */
  export let careers = [];

  /** @type {string} */
  export let error = "";
  console.log(error);


  $: if (error === "expired") {
    error = "Your session has expired. Please login again";
  } else if (error === "invalid credentials") {
    error = "Invalid email or password";
  } else if (error === "error generating token") {
    error = "An error occurred while generating your token. Please try again";
  }

  $: if (error) {
    setTimeout(() => {
      error = "";
      window.history.replaceState({}, "", "/login");
    }, 4000);
  }
</script>

<MainLayout
  title="LogIn to the MockU System - Mock University"
  description="Login to the MockU system to access your account"
  {careers}
>
  <div class="w-full lg:grid lg:grid-cols-2 min-h-scn">
    <div class="flex items-center justify-center py-12">
      <form
        method="post"
        action="/login_post"
        class="mx-auto grid w-[350px] gap-6"
        enctype="multipart/form-data"
      >
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Login</h1>
          <p class="text-balance text-muted-foreground">
            Enter your email below to login to your account
          </p>
        </div>
        <div class="grid gap-4">
          <div class="grid gap-2">
            <Label for="email">Email</Label>
            <Input
              id="email"
              type="email"
              placeholder="m@example.com"
              required
              name="email"
            />
          </div>
          <div class="grid gap-2">
            <div class="flex items-center">
              <Label for="password">Password</Label>
              <a href="##" class="ml-auto inline-block text-sm underline">
                Forgot your password?
              </a>
            </div>
            <Input id="password" type="password" name="password" required />
          </div>
          {#if error}
            <p class="text-red-500 text-sm">{error}</p>
          {/if}
          <Button type="submit" class="w-full">Login</Button>
          <Button variant="outline" class="w-full">Login with Google</Button>
        </div>
      </form>
    </div>
    <div class="hidden bg-muted lg:block">
      <img
        src="/images/placeholder.svg"
        alt="placeholder"
        width="1920"
        height="1080"
        class="h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
      />
    </div>
  </div>
</MainLayout>
