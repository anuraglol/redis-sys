<script lang="ts">
	import "./layout.css";
	import favicon from "$lib/assets/favicon.svg";
	import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
	import { SvelteQueryDevtools } from "@tanstack/svelte-query-devtools";
	import { Toaster } from "$lib/components/ui/sonner";
	import HomeSidebar from "$lib/components/home/home-sidebar.svelte";

	const queryClient = new QueryClient();
	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<QueryClientProvider client={queryClient}>
	<Toaster />
	<div class="flex h-screen w-screen bg-background text-foreground overflow-hidden font-sans antialiased">
		<HomeSidebar />
		<main class="flex-1 flex flex-col overflow-hidden relative min-w-0">
			{@render children()}
		</main>
	</div>
	<SvelteQueryDevtools />
</QueryClientProvider>
