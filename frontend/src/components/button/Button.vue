<script setup lang="ts">
import { ref } from 'vue'
import { MagnifyingGlassIcon } from '@heroicons/vue/24/outline'
interface ButtonProps {
  label: string
}

const props = defineProps<ButtonProps>()

const emit = defineEmits(['clicked'])

const handleClick = (event: Event) => {
  emit('clicked')
}
</script>

<template>
  <button class="Button" @click="handleClick">{{ props.label }}</button>
</template>

<style lang="postcss" scoped>
.Button {
  display: inline-flex;
  background-color: #8948db;
  align-items: center;
  cursor: pointer;
  justify-content: center;
  height: 40px;
  color: #fff;
  padding: 0.5rem 1.5rem;
  border-radius: 8px;
  border: 0;
  font-family: 'Inter', sans-serif;
  font-size: 1rem;
  font-weight: 500;
  box-sizing: border-box;
  appearance: none;
  white-space: nowrap;
  -webkit-appearance: none;
  width: 100%;
}
</style>
