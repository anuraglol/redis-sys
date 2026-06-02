<script lang="ts">
    import SidebarHeader from "./sidebar-header.svelte";
    import SidebarTabs from "./sidebar-tabs.svelte";
    import SidebarKeyList from "./sidebar-key-list.svelte";
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
    <SidebarHeader {statsQuery} />
    <SidebarTabs {activeTab} {onTabChange} />
    <div class="flex-1 overflow-y-auto p-4 space-y-4">
        <SidebarKeyList {statsQuery} {onTabChange} />
    </div>
    <SidebarStatsFooter {statsQuery} />
</aside>
