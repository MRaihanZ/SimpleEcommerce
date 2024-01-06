<template>
	<div>
		<div v-for="p in products">
			<div class="row border border-1 py-2 bg-white">
				<div class="col">
					<div class="row">
						<div class="col-md-1">
							<img :src="p.data.image" alt="" class="img-thumbnail m-auto" />
						</div>
						<div class="col-md-10">
							<p>{{ p.data.title }}</p>
							<p>Rp. {{ p.data.price }}k</p>
							<p>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="16"
									height="16"
									fill="currentColor"
									class="bi bi-star-fill mb-1"
									viewBox="0 0 16 16"
								>
									<path
										d="M3.612 15.443c-.386.198-.824-.149-.746-.592l.83-4.73L.173 6.765c-.329-.314-.158-.888.283-.95l4.898-.696L7.538.792c.197-.39.73-.39.927 0l2.184 4.327 4.898.696c.441.062.612.636.282.95l-3.522 3.356.83 4.73c.078.443-.36.79-.746.592L8 13.187l-4.389 2.256z"
									/>
								</svg>
								{{ p.data.rating }}
							</p>
						</div>
						<div class="col-md-1">
							<button
								class="btn btn-primary btn-cart"
								@click="
									addToCart(
										p.data.id,
										p.data.title,
										p.data.price,
										p.data.image,
										p.data.rating
									)
								"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="16"
									height="16"
									fill="currentColor"
									class="bi bi-cart-plus"
									viewBox="0 0 16 16"
								>
									<path
										d="M9 5.5a.5.5 0 0 0-1 0V7H6.5a.5.5 0 0 0 0 1H8v1.5a.5.5 0 0 0 1 0V8h1.5a.5.5 0 0 0 0-1H9z"
									/>
									<path
										d="M.5 1a.5.5 0 0 0 0 1h1.11l.401 1.607 1.498 7.985A.5.5 0 0 0 4 12h1a2 2 0 1 0 0 4 2 2 0 0 0 0-4h7a2 2 0 1 0 0 4 2 2 0 0 0 0-4h1a.5.5 0 0 0 .491-.408l1.5-8A.5.5 0 0 0 14.5 3H2.89l-.405-1.621A.5.5 0 0 0 2 1zm3.915 10L3.102 4h10.796l-1.313 7h-8.17zM6 14a1 1 0 1 1-2 0 1 1 0 0 1 2 0m7 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0"
									/>
								</svg>
							</button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
const { data: products } = await useFetch("http://localhost:8080/api/products");

async function addToCart(id, title, price, image, rating) {
	const requestBody = {
		id_product: id,
		title: title,
		price: price,
		image: image,
		rating: rating,
		quantity: 1,
	};

	await $fetch("http://localhost:8080/api/cart", {
		method: "POST",
		headers: {
			"client-platform": "json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify(requestBody),
	});
}
</script>

<style scoped>
.btn-cart {
	height: 100%;
	width: 100%;
}
</style>
