<script>
  import "$lib/styles";
</script>

<div class="font-sans content dark:text-gray-100">
  <slot />
</div>

<style>
  :global(body) {
    --uno: "dark:bg-dark-900";
  }
</style>
