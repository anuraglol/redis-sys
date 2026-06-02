<script lang="ts">
    import { Database, RefreshCw } from "@lucide/svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import { queryClient } from "$lib/queries/query-client";
    import type { StatsQuery } from "$lib/queries/queries";

    let { statsQuery }: { statsQuery: StatsQuery } = $props();
</script>

<div class="p-6 border-b border-border flex items-center justify-between">
    <div class="flex items-center gap-3">
        <div
            class="h-9 w-9 rounded-xl bg-primary flex items-center justify-center shadow-md shadow-primary/20"
        >
            <Database class="h-5 w-5 text-primary-foreground" />
        </div>
        <div>
            <h1 class="text-sm font-bold tracking-tight uppercase">Kotori</h1>
            <p class="text-xs text-muted-foreground font-mono">v1.0.0-beta</p>
        </div>
    </div>
    <Button
        variant="ghost"
        size="icon"
        class="h-8 w-8 rounded-lg"
        onclick={() => queryClient.invalidateQueries({ queryKey: ["stats"] })}
    >
        <RefreshCw
            class="h-4 w-4 {statsQuery.isFetching ? 'animate-spin' : ''}"
        />
    </Button>
</div>
