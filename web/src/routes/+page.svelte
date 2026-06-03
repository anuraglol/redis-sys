<script lang="ts">
    import HomeSidebar from "$lib/components/home/home-sidebar.svelte";
    import HomeCliView from "$lib/components/home/home-cli-view.svelte";
    import TableView from "$lib/components/home/table-view.svelte";

    let activeTab = $state("data");
</script>

<div
    class="flex h-screen w-screen bg-background text-foreground overflow-hidden font-sans antialiased selection:bg-primary/10"
>
    <HomeSidebar {activeTab} onTabChange={(t) => (activeTab = t)} />
    <main class="flex-1 bg-background flex flex-col overflow-hidden relative">
        {#if activeTab === "cli"}
            <HomeCliView />
        {/if}
        {#if activeTab === "data"}
            <TableView />
        {/if}
    </main>
</div>
