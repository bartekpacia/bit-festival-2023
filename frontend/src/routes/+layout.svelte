<script>
	import "./styles.css";
	import Header from "$lib/components/Header.svelte";
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		display: flex;
		width: 100%;
		flex-grow: 1;
		align-items: center;
	}
</style>
