<template>
  <div
    :class="['pop-up', { active: isChangePassword }]"
    @click.stop="$emit('closeChangePassword')"
  >
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="confirm__body">
          <div class="confirm__title">Açar sözi üýtgetmek</div>
          <div class="communication_form form__confirm">
            <div class="form__box">
              <span>Acar soz</span>
              <input type="text" placeholder="Açar soz" />
            </div>
            <div class="form__box">
              <span>Acar sozi tassykla</span>
              <input type="text" placeholder="Acar sozi tassykla" />
            </div>
          </div>
          <div class="popup-confirm__text">
            Siziň açar sözüňiz 5-20 uzynlygynda, harp ýa-da sandan ybarat
            bolmalydyr
          </div>
          <div class="popup-confirm__btn">
            <button class="product__add-btn confirm__btn">
              <h4>Ýatda sakla</h4>
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
    isChangePassword: {
      type: Boolean,
      default: false,
    },
  },
}
</script>
