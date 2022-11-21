export default {
  methods: {
    translationProductName(translations) {
      let translation = translations?.filter(
        (translation) => translation[this.$i18n.locale]
      )
      return translation[0][this.$i18n.locale]?.name
    },
    translationProductDescription(translations) {
      let translation = translations?.filter(
        (translation) => translation[this.$i18n.locale]
      )
      return translation[0][this.$i18n.locale]?.description
    },
  },
}
