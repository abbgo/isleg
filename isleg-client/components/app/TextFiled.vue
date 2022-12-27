<template>
  <div class="input" :class="{ error: error }">
    <div v-if="label" class="input__label">
      <span>{{ label }}</span>
    </div>
    <div class="input__body">
      <div v-if="prepenedIcon" class="prepend__icon">
        <img :src="require(`@/assets/img/${prepenedIcon}`)" alt="" />
      </div>
      <div class="input__input">
        <input
          :type="type"
          v-model="modelValue"
          :disabled="disabled"
          @input="enforcePhoneFormat"
          :placeholder="placeholder"
        />
      </div>

      <div v-if="appendIcon" class="append__icon">
        <img :src="require(`@/assets/img/${appendIcon}`)" alt="" />
      </div>
    </div>
    <div v-if="error" class="input__error">
      <span>{{ errorText }}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: [String, Number],
      default: null,
    },
    label: {
      type: String,
      default: '',
    },
    prepenedIcon: {
      type: String,
      default: '',
    },
    appendIcon: {
      type: String,
      default: '',
    },
    placeholder: {
      type: String,
      default: '',
    },
    type: {
      type: String,
      default: 'text',
    },
    error: {
      type: Boolean,
      default: false,
    },
    errorText: {
      type: String,
      default: '',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      modelValue: this.value,
    }
  },
  watch: {
    value: function (val) {
      this.modelValue = val
    },
  },
  methods: {
    enforcePhoneFormat() {
      if (this.type == 'tel') {
        let x = this.modelValue
          .replace(/\D/g, '')
          .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
        this.modelValue = !x[2]
          ? '+9936'
          : '+9936' +
            (x[3] ? x[3] : '') +
            (x[4] ? x[4] : '') +
            (x[5] ? x[5] : '') +
            (x[6] ? x[6] : '')
        this.$emit('updateValue', this.modelValue)
      } else {
        this.$emit('updateValue', this.modelValue)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.input {
  width: 50%;
  position: relative;
  padding-bottom: 12px;
  padding-left: 20px;
  padding-right: 20px;
  margin-bottom: 10px;
  @media (max-width: 950px) {
    width: 100%;
  }
  // @media (max-width: 435px) {
  //   padding-left: 0px;
  //   padding-right: 0px;
  // }
  &__label {
    margin: 5px 2px;
    span {
      font-family: TTNormsPro;
      font-style: normal;
      font-weight: 800;
      font-size: 16px;
      line-height: 120%;
      text-align: center;
      color: #1b3254;
    }
  }
  &__body {
    display: flex;
    align-items: center;
    background: #eee;
    border-radius: 5px;
  }
  &__input {
    flex: 1 1 auto;
    padding: 15px 8px;
    input {
      width: 100%;
      font-size: 16px;
      background: transparent;
    }
  }
  &__error {
    position: absolute;
    top: 62px;
    padding: 2px;
    span {
      font-family: TTNormsPro;
      font-style: normal;
      font-weight: 500;
      font-size: 12px;
      line-height: 120%;
      color: #aa0000;
    }
  }

  .prepend__icon {
    width: 16px;
    height: 16px;
    padding-left: 5px;
    img {
      width: 100%;
      height: 100%;
      object-fit: contain;
      object-position: center;
    }
  }
  .append__icon {
    width: 16px;
    height: 16px;
    padding-right: 5px;
    img {
      width: 100%;
      height: 100%;
      object-fit: contain;
      object-position: center;
    }
  }
}

.input.error {
  animation: 0.2s invalid forwards;
}
.input.error .input__label span {
  color: #aa0000;
}
.input.error .input__body {
  border: 1px solid #aa0000;
}
.input.error .input__body input::placeholder {
  color: #aa0000;
  opacity: 0.6;
}

@keyframes invalid {
  0% {
    transform: translateX(0px);
  }
  25% {
    transform: translateX(5px);
  }
  50% {
    transform: translateX(0px);
  }
  75% {
    transform: translateX(-5px);
  }
  100% {
    transform: translateX(0px);
  }
}
</style>
