<template>
	<div>
		<div v-for="p in products">
			<div class="row border border-1 py-2 bg-white">
				<div class="col">
					<div class="row">
						<div class="col-md-1">
							<img :src="p.image" alt="" class="img-thumbnail m-auto" />
						</div>
						<div class="col-md-10">
							<p>{{ p.title }}</p>
							<p>Rp. {{ p.price }}k</p>
							<span>
								<button
									class="btn btn-primary me-2"
									@click="QuantityCart(p.id_product, p.quantity, 'decrease')"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="currentColor"
										class="bi bi-dash"
										viewBox="0 0 16 16"
									>
										<path
											d="M4 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 4 8"
										/>
									</svg>
								</button>
								<input
									type="number"
									class="quantity-input d-inline-block form-control"
									:value="p.quantity"
								/>

								<button
									class="btn btn-primary ms-2"
									@click="QuantityCart(p.id_product, p.quantity, 'increase')"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="currentColor"
										class="bi bi-plus"
										viewBox="0 0 16 16"
									>
										<path
											d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4"
										/>
									</svg>
								</button>
							</span>
						</div>
						<div class="col-md-1">
							<button
								class="btn btn-danger btn-cart"
								@click="deleteFromCart(p.id_product)"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="16"
									height="16"
									fill="currentColor"
									class="bi bi-cart-dash"
									viewBox="0 0 16 16"
								>
									<path d="M6.5 7a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1z" />
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
const { data: products } = await useFetch("http://localhost:8080/api/carts");

async function deleteFromCart(id) {
	await $fetch(`http://localhost:8080/api/cart/${id}`, {
		method: "DELETE",
	});
	await refreshNuxtData();
}

async function QuantityCart(id, currentQuantity, action) {
	if (action === "increase") {
		let quantity = currentQuantity + 1;
		await $fetch(`http://localhost:8080/api/cart/${id}/${quantity}`, {
			method: "PATCH",
		});
		await refreshNuxtData();
	} else if (action === "decrease") {
		let quantity = currentQuantity - 1;
		await $fetch(`http://localhost:8080/api/cart/${id}/${quantity}`, {
			method: "PATCH",
		});
		await refreshNuxtData();
	}
}
</script>

<style scoped>
.btn-cart {
	height: 100%;
	width: 100%;
}

.quantity-input {
	width: 4rem;
}
</style>
