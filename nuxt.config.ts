// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	plugins: ["~/plugins/axios"],
	devtools: { enabled: true },
	css: [
		"~/node_modules/bootstrap/dist/css/bootstrap.min.css",
		"~/assets/css/components/navbar.css",
		"~/assets/css/global.css",
	],
});
