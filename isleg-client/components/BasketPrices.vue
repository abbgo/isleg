<template>
	<div class="basket__prices">
		<price-items
			:count="productCount"
			:productsChek="productsChek"
			:totalPrice="totalPrice"
		></price-items>
		<total-price
			:totalPrice="totalPrice"
			@payment="$emit('openPayment')"
		></total-price>
	</div>
</template>

<script>
export default {
	props: {
		productCount: {
			type: Number,
			default: () => 0,
		},
		productsChek: {
			type: Object,
			default: () => {},
		},
		totalPrice: {
			type: Number,
			default: () => 0,
		},
	},
}
</script>
