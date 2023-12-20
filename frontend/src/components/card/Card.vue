<script setup lang="ts">
interface CardProps {
  title: string
  value: string | number
}

const props = defineProps<CardProps>()
</script>

<template>
  <div className="Card">
    <p className="CardTitle">{{ props.title }}</p>
    <h4 className="CardValue">{{ props.value }}</h4>
  </div>
</template>

<style lang="postcss" scoped>
.Card {
  @apply w-full border border-[#ECEBF5] p-5 rounded-xl md:mb-0 mb-6;
  box-shadow: 0px 1px 12px 0px #15295208;

  .CardTitle {
    @apply text-[#7676A3] text-sm mb-3;
  }

  .CardValue {
    @apply text-[##0E0842] text-2xl;
  }
}
</style>
