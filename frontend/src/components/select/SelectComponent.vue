<!-- import styled from "styled-components";

interface OptionProps {
  label: string;
  value: string | number;
}


const SelectWrapper = styled.div`
  position: relative;
`;

const SelectField = styled.select`
 
`;

const OptionField = styled.option`
  appearance: none;
  height: 40px;
  width: 100%;
  padding: 6px 10px;
  box-sizing: border-box;
  border: 1px solid #dee5e7;
  color: #6f767e;
  border-radius: 5px;
  transition: 0.3s ease-in;
`;

const Select: React.FC<SelectProps> = ({
  handleChange,
  customStyles,
  name,
  options,
  value,
}) => {
  return (
    <SelectWrapper style={customStyles}>
      <SelectField
        name={name}
        id={name}
        onChange={handleChange}
        defaultValue={value}
      >
        {options.map((option: OptionProps, i: number) => (
          <OptionField value={option.value} key={i}>
            {option.label}
          </OptionField>
        ))}
      </SelectField>
    </SelectWrapper>
  );
};

export default Select; -->

<script setup lang="ts">
interface OptionProps {
  label: string
  value: string | number
}

interface SelectProps {
  customStyles?: Record<string, string>
  value: string
  modelValue: string | number
  options: OptionProps[]
}

const props = defineProps<SelectProps>()

const emit = defineEmits(['update:modelValue'])

const updateValue = (event) => {
  emit('update:modelValue', event.target.value)
}
</script>

<template>
  <div class="SelectWrapper" :style="props.customStyles">
    <!-- {icon && icon} -->
    <!-- <input class="SelectField" :value="modelValue" @input="updateValue" />-->

    <select
      class="SelectField"
      placeholder="Sélectionnez un diplôme"
      :value="value"
      @change="updateValue"
      expanded
    >
      <option v-for="option in options" :value="option.value" :key="option.value">
        {{ option.label }}
      </option>
    </select>
  </div>
</template>

<style lang="postcss" scoped>
.SelectWrapper {
  @apply w-full;

  position: relative;

  .SelectField {
    appearance: none;

    height: 40px;
    width: 100%;
    padding: 6px 25px 6px 10px;
    box-sizing: border-box;
    border: 1px solid #dee5e7;
    color: #6f767e;
    border-radius: 5px;
    transition: 0.3s ease-in;
    outline: none;
    &:focus {
      border: 1px solid #999dff;
    }
    &::placeholder {
      color: #aab1b9;
    }
  }

  .OptionField {
    appearance: none;
    height: 40px;
    width: 100%;
    padding: 6px 10px;
    box-sizing: border-box;
    border: 1px solid #dee5e7;
    color: #6f767e;
    border-radius: 5px;
    transition: 0.3s ease-in;
  }
}
</style>
