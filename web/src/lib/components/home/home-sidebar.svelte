<script lang="ts">
    import SidebarTabs from "./sidebar-tabs.svelte";
    import SidebarStatsFooter from "./sidebar-stats-footer.svelte";
    import { getStatsQuery } from "$lib/queries/queries";

    let {
        activeTab,
        onTabChange,
    }: {
        activeTab: string;
        onTabChange: (tab: string) => void;
    } = $props();

    const statsQuery = getStatsQuery();
</script>

<aside
    class="w-80 border-r border-border bg-card/50 backdrop-blur-md flex flex-col z-10 shrink-0"
>
    <SidebarTabs {activeTab} {onTabChange} />
    <SidebarStatsFooter {statsQuery} />
</aside>
