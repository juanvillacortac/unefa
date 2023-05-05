<script lang="ts">
  import { applyAction, enhance } from "$app/forms";
  import { pb } from "$lib/pocketbase";
  import logoSrc from "$lib/assets/logo.png";
  import { slide } from "svelte/transition";
  export let form;
  $: console.log(form);
</script>

<form
  method="POST"
  class="p-4 flex flex-col space-y-8 w-full p-2 sm:w-40% lg:w-18% items-center"
  use:enhance={() => {
    return async ({ result }) => {
      console.log(result);
      pb.authStore.loadFromCookie(document.cookie);
      applyAction(result);
    };
  }}
>
  <img
    src={logoSrc}
    alt="UNEFA"
    class="flex w-60% dark:(filter-brightness-0 invert)"
  />
  <div class="flex-col space-y-2 flex flex-centered">
    <div data-flex data-p-4 data-bg="light-700 dark:dark-800" data-rounded-full>
      <un-i-carbon-user data data-text-xl />
    </div>
    <h1 class="text-xl text-center flex-inline space-x-2 flex-centered">
      <span>Inicio de sesión</span>
    </h1>
    {#if form?.error}
      <div
        class="text-sm text-red-400 text-center"
        in:slide={{ duration: 200 }}
      >
        {form.error}
      </div>
    {/if}
  </div>
  <div class="flex flex-col space-y-2 w-full">
    <label class="flex flex-col space-y-2">
      <span class="text-xs font-bold">Correo</span>
      <input
        required
        type="email"
        name="email"
        placeholder="Ej. mariana@gmail.com"
        class="bg-light-700 p-2 rounded text-gray-700 w-full flex dark:bg-dark-700 dark:text-white"
      />
    </label>
    <label class="flex flex-col space-y-2">
      <span class="text-xs font-bold">Contraseña</span>
      <input
        required
        type="password"
        name="password"
        class="bg-light-700 p-2 rounded text-gray-700 w-full flex dark:bg-dark-700 dark:text-white"
      />
    </label>
  </div>
  <button
    class="btn btn-primary w-full flex-inline space-x-2 flex-centered group"
  >
    <span>Iniciar sesión</span>
    <span
      data-i-carbon-arrow-right
      class="transform group-hover-translate-x-0.2rem"
    />
  </button>
</form>
