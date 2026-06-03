<script lang="ts">
    import {
        Clock,
        Layers,
        RefreshCw,
        Trash2,
        KeyRound,
        Plus,
    } from "@lucide/svelte";
    import * as Card from "$lib/components/ui/card";
    import Button from "$lib/components/ui/button/button.svelte";
    import Badge from "$lib/components/ui/badge/badge.svelte";
    import { homeState } from "$lib/state/home.svelte";
    import { getValueQuery, getTtlQuery } from "$lib/queries/queries";
    import {
        createIncrMutation,
        createDeleteMutation,
    } from "$lib/queries/mutations";

    const valueQuery = getValueQuery();
    const ttlQuery = getTtlQuery();
    const incrMutation = createIncrMutation();
    const deleteMutation = createDeleteMutation();
</script>

<Card.Root class="xl:col-span-2 shadow-sm border border-border bg-card">
    <Card.Header class="pb-3 border-b border-border/60">
        <Card.Title class="text-base font-bold tracking-tight"
            >Active Key Analyzer</Card.Title
        >
        <Card.Description class="text-xs"
            >Inspect runtime cluster objects mapping execution pipelines</Card.Description
        >
    </Card.Header>
    <Card.Content class="pt-6">
        {#if !homeState.selectedKey}
            <div
                class="flex flex-col items-center justify-center py-20 text-center border border-dashed border-border rounded-xl bg-muted/20"
            >
                <div
                    class="h-10 w-10 rounded-xl bg-muted flex items-center justify-center mb-4"
                >
                    <KeyRound class="h-5 w-5 text-muted-foreground" />
                </div>
                <h3 class="text-sm font-semibold">No Target Initialized</h3>
                <p class="text-xs text-muted-foreground mt-1 max-w-xs">
                    Select an operational telemetry key variant from the index
                    tracker boundary interface.
                </p>
            </div>
        {:else}
            <div class="space-y-6">
                <div
                    class="flex flex-wrap items-center justify-between gap-4 p-4 bg-muted/40 border border-border rounded-xl"
                >
                    <div class="space-y-1">
                        <span
                            class="text-[10px] font-bold uppercase tracking-wider text-muted-foreground font-mono block"
                            >Selected Object Identity</span
                        >
                        <span
                            class="font-mono text-sm font-semibold text-primary"
                            >{homeState.selectedKey}</span
                        >
                    </div>
                    <div class="flex items-center gap-2">
                        <Button
                            variant="outline"
                            size="sm"
                            class="h-8 text-xs font-medium"
                            onclick={() =>
                                incrMutation.mutate(homeState.selectedKey)}
                            disabled={incrMutation.isPending}
                        >
                            {#if incrMutation.isPending}
                                <RefreshCw
                                    class="w-3.5 h-3.5 mr-1.5 animate-spin"
                                />
                            {:else}
                                <Plus class="w-3.5 h-3.5 mr-1.5" />
                            {/if}
                            Incr
                        </Button>
                        <Button
                            variant="destructive"
                            size="sm"
                            class="h-8 text-xs font-medium"
                            onclick={() =>
                                deleteMutation.mutate(homeState.selectedKey)}
                            disabled={deleteMutation.isPending}
                        >
                            <Trash2 class="w-3.5 h-3.5 mr-1.5" />
                            Drop Key
                        </Button>
                    </div>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div
                        class="p-4 border border-border rounded-xl bg-card/50 space-y-2"
                    >
                        <div
                            class="flex items-center gap-2 text-xs text-muted-foreground font-medium"
                        >
                            <Layers class="w-3.5 h-3.5 text-primary" /> Array Bytes
                            Payload
                        </div>
                        {#if valueQuery.isLoading}
                            <div
                                class="h-6 bg-muted rounded animate-pulse w-3/4"
                            ></div>
                        {:else if valueQuery.error}
                            <span class="text-xs text-destructive"
                                >Data link corrupted</span
                            >
                        {:else if valueQuery.data?.isNil}
                            <span
                                class="font-mono text-xs text-muted-foreground italic"
                                >nil</span
                            >
                        {:else}
                            <div
                                class="font-mono text-sm font-bold bg-muted/60 p-2.5 rounded-lg border border-border break-all max-h-40 overflow-y-auto"
                            >
                                {JSON.stringify(valueQuery.data?.value)}
                            </div>
                        {/if}
                    </div>

                    <div
                        class="p-4 border border-border rounded-xl bg-card/50 space-y-2"
                    >
                        <div
                            class="flex items-center gap-2 text-xs text-muted-foreground font-medium"
                        >
                            <Clock class="w-3.5 h-3.5 text-primary" /> Temporal TTL
                            Matrix
                        </div>
                        {#if ttlQuery.isLoading}
                            <div
                                class="h-6 bg-muted rounded animate-pulse w-1/2"
                            ></div>
                        {:else}
                            <div class="flex items-center gap-2">
                                {#if ttlQuery.data?.ttl === -1}
                                    <Badge
                                        variant="secondary"
                                        class="bg-blue-500/10 text-blue-500 border border-blue-500/20 font-mono"
                                        >Persistent (-1)</Badge
                                    >
                                {:else if ttlQuery.data?.ttl === -2 || !ttlQuery.data}
                                    <Badge
                                        variant="destructive"
                                        class="font-mono"
                                        >Expired / Void (-2)</Badge
                                    >
                                {:else}
                                    <Badge
                                        variant="outline"
                                        class="bg-amber-500/10 text-amber-500 border border-amber-500/20 font-mono animate-pulse"
                                    >
                                        {ttlQuery.data.ttl} seconds
                                    </Badge>
                                {/if}
                            </div>
                        {/if}
                    </div>
                </div>
            </div>
        {/if}
    </Card.Content>
</Card.Root>
