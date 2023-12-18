<script setup lang="ts">
import { ref } from 'vue'
interface InputProps {
  customStyles?: Record<string, string>
  customClass?: string
  modelValue: string
  icon?: any
  type?: string
  placeholder?: string
}

const props = defineProps<InputProps>()

const emit = defineEmits(['update:modelValue'])

const updateValue = (event) => {
  emit('update:modelValue', event.target.value)
}

const type = ref(props.type || 'text')
</script>

<template>
  <div class="InputWrapper" :style="props.customStyles">
    <!-- {icon && icon} -->
    <input
      class="InputField"
      :value="modelValue"
      :placeholder="props.placeholder"
      :type="type"
      @input="updateValue"
    />
  </div>
</template>

<style lang="postcss" scoped>
.InputWrapper {
  @apply w-full;

  position: relative;

  .InputField {
    @apply appearance-none w-full rounded-md box-border;
    height: 40px;
    padding: 6px 10px;
    border: 1px solid #dee5e7;
    color: #6f767e;
    transition: 0.3s ease-in;
    outline: none;
    font-size: 13px;
    &.icon {
      padding-left: 30px;
    }
    &:focus {
      border: 1px solid #999dff;
    }
    &::placeholder {
      color: #aab1b9;
    }
  }

  svg {
    color: #999dff;
    top: 50%;
    transform: translateY(-50%);
    left: 10px;
    position: absolute;
  }
}
</style>
