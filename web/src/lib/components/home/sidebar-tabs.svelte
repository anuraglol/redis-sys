<script lang="ts">
    import { Layers, Terminal } from "@lucide/svelte";

    let {
        activeTab,
        onTabChange,
    }: {
        activeTab: string;
        onTabChange: (tab: string) => void;
    } = $props();
</script>

<div class="p-4 border-b border-border bg-muted/40 h-16">
    <div class="flex p-1 bg-muted rounded-lg gap-1">
        <button
            class="flex-1 py-1.5 text-xs font-medium rounded-md transition-all flex items-center justify-center gap-2 {activeTab ===
            'data'
                ? 'bg-background shadow-sm text-foreground'
                : 'text-muted-foreground hover:text-foreground'}"
            onclick={() => onTabChange("data")}
        >
            <Layers class="w-3.5 h-3.5" /> Table
        </button>
        <button
            class="flex-1 py-1.5 text-xs font-medium rounded-md transition-all flex items-center justify-center gap-2 {activeTab ===
            'cli'
                ? 'bg-background shadow-sm text-foreground'
                : 'text-muted-foreground hover:text-foreground'}"
            onclick={() => onTabChange("cli")}
        >
            <Terminal class="w-3.5 h-3.5" /> Console
        </button>
    </div>
</div>
