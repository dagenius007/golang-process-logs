<script setup lang="ts">
import { ref } from 'vue'
import { RectangleStackIcon, Bars3BottomRightIcon, XMarkIcon } from '@heroicons/vue/24/outline'

function handleToggle() {
  if (document) {
    const sidebar = document.getElementById('sidebar') as HTMLElement
    if (sidebar.classList.contains('-translate-x-60')) {
      // -translate-x-60
      sidebar.classList.remove('-translate-x-60')
      sidebar.classList.add('translate-x-0')
    } else {
      sidebar.classList.remove('translate-x-0')
      sidebar.classList.add('-translate-x-60')
    }
  }
}
</script>

<template>
  <div class="Layout">
    <div class="flex w-full">
      <aside id="sidebar" aria-label="Sidebar" class="SideNav md:translate-x-0 -translate-x-60">
        <XMarkIcon
          class="h-7 w-7 absolute top-4 right-6 text-[#4F4999] md:hidden"
          @click="handleToggle"
        />
        <p class="Logo">Logo</p>

        <nav class="NavList">
          <div class="Nav">
            <RectangleStackIcon class="h-5 w-5 mr-2 text-[#4F4999]" />
            <p>Process</p>
          </div>
        </nav>
      </aside>
      <div class="Main">
        <div class="MainHeader">
          <Bars3BottomRightIcon class="h-6 w-6 mr-2 text-[#4F4999]" @click="handleToggle" />

          <p>{{ $route.name }}</p>
        </div>
        <section class="Section">
          <RouterView> </RouterView>
        </section>
      </div>
    </div>
  </div>
</template>

<style lang="postcss" scoped>
.Layout {
  @apply w-screen overflow-x-hidden;
  .SideNav {
    @apply fixed h-screen z-40 w-60 border border-[#EAECF0] bg-[#FAFBFF] px-6 py-10;
    transition: all 0.5s ease;
  }
  .Logo {
    @apply text-3xl mb-12;
  }
  .Main {
    @apply md:w-[calc(100%-240px)] ml-auto w-full;
  }
  .MainHeader {
    @apply md:px-12 px-4 py-7 font-medium text-[#0E0842] border-b-[1px] text-lg flex;
    & > svg {
      @apply md:hidden mr-6;
    }
  }
  .Section {
    @apply md:p-12 p-6;
  }
  .Nav {
    @apply text-sm text-[#4F4999] mb-6 flex;
    .NavIcon {
      @apply w-6 h-6;
    }
  }
}
</style>
