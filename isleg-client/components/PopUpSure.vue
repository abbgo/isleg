<template>
  <div :class="['pop-up', { active: isSure }]" @click.stop="$emit('close')">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__sure">
          <div class="pop-up__text">
            Siz hakykatdanam harydy sebetden aýyrmak isleýäňizmi?
          </div>
          <div class="pop-up__action">
            <button class="pop-up__btn-cancel btn" @click.stop="$emit('close')">
              Ýok</button
            ><button class="pop-up__btn-confirm btn" @click="$emit('confirm')">
              Hawa
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isSure: {
      type: Boolean,
      default: false,
    },
  },
}
</script>
