export default function (context) {
  console.log(context.store)
  context.store.dispatch('ui/initAuth')
}
