<script>
  import { applyAction, enhance } from "$app/forms";
  import { pb } from "$lib/pocketbase";
</script>

<form
  method="POST"
  action="/logout"
  use:enhance={() => {
    return async ({ result }) => {
      pb.authStore.clear();
      await applyAction(result);
    };
  }}
>
  <slot />
</form>
