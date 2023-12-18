<script setup lang="ts"></script>

<template>
  <div class="Layout">
    <p>{{ $route.name }}</p>
    <RouterView></RouterView>
  </div>
</template>

<style lang="postcss" scoped>
.Layout {
  @apply p-32;
}
</style>
