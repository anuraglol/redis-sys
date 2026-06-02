<script lang="ts">
	import HomeSidebar from "$lib/components/home/home-sidebar.svelte";
	import HomeDataView from "$lib/components/home/home-data-view.svelte";
	import HomeCliView from "$lib/components/home/home-cli-view.svelte";

	let activeTab = $state("data");
</script>

<div
	class="flex h-screen w-screen bg-background text-foreground overflow-hidden font-sans antialiased selection:bg-primary/10"
>
	<HomeSidebar {activeTab} onTabChange={(t) => (activeTab = t)} />
	<main class="flex-1 bg-background flex flex-col overflow-hidden relative">
		{#if activeTab === "data"}
			<HomeDataView />
		{:else if activeTab === "cli"}
			<HomeCliView />
		{/if}
	</main>
</div>
