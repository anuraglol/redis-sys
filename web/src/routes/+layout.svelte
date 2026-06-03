<script lang="ts">
    import "./layout.css";
    import favicon from "$lib/assets/favicon.svg";
    import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
    import { SvelteQueryDevtools } from "@tanstack/svelte-query-devtools";
    import { Toaster } from "$lib/components/ui/sonner";

    const queryCient = new QueryClient();
    let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<QueryClientProvider client={queryCient}>
    <Toaster />
    {@render children()}
    <SvelteQueryDevtools />
</QueryClientProvider>
