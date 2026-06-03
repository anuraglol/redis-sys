<script lang="ts">
    import { getAllQuery } from "$lib/queries/queries";
    import { Loader2, AlertCircle, Inbox, RefreshCw } from "@lucide/svelte";
    import * as Table from "$lib/components/ui/table";
    import EditEntryDialog from "./edit-entry-dialog.svelte";
    import DeleteEntryDialog from "./delete-entry-dialog.svelte";
    import CreateEntryDialog from "./create-entry-dialog.svelte";
    import FlushAllDialog from "./flush-all-dialog.svelte";
    import Seed from "./seed.svelte";

    const getAll = getAllQuery();

    const entries = $derived(
        getAll.data
            ? Object.entries(getAll.data.data).map(([key, item]) => ({
                  key,
                  ...item,
              }))
            : [],
    );
</script>

<header
    class="h-16 border-b border-border px-8 flex items-center justify-between bg-card/30 backdrop-blur-md shrink-0"
>
    <div class="flex items-center gap-3">
        <h2 class="text-sm font-semibold tracking-tight">Data View Tada</h2>
    </div>
    <div class="flex gap-6 items-center">
        <span class="relative flex size-2.5">
            {#if getAll.isFetching}
                <span
                    class="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-200 opacity-75"
                ></span>
            {/if}
            <span
                class="relative inline-flex size-2.5 rounded-full bg-green-300"
            ></span>
        </span>
        <Seed />
        <FlushAllDialog />
        <CreateEntryDialog />
        <span class="text-xs text-zinc-500">
            {#if getAll.isSuccess}
                {entries.length} {entries.length === 1 ? "entry" : "entries"}
            {/if}
        </span>
    </div>
</header>

<div
    class="flex-1 bg-black text-zinc-200 p-6 overflow-y-auto flex flex-col selection:bg-zinc-800"
>
    {#if getAll.isPending}
        <div class="flex items-center justify-center gap-2 text-zinc-500 py-12">
            <Loader2 class="size-4 animate-spin" />
            <span class="text-sm">loading entries…</span>
        </div>
    {:else if getAll.isError}
        <div
            class="flex flex-col items-center justify-center gap-2 text-red-400 py-12"
        >
            <AlertCircle class="size-5" />
            <p class="text-sm">failed to load data</p>
            <p class="text-xs text-zinc-500">
                {getAll.error?.message ?? "unknown error"}
            </p>
        </div>
    {:else if entries.length === 0}
        <div
            class="flex flex-col items-center justify-center gap-2 text-zinc-500 py-12"
        >
            <Inbox class="size-5" />
            <p class="text-sm">no entries yet</p>
            <p class="text-xs">add a key to get started</p>
        </div>
    {:else}
        <div class="rounded-md border border-border">
            <Table.Root class="text-zinc-200 border-zinc-800">
                <Table.Header class="border-zinc-800 hover:bg-transparent">
                    <Table.Row class="hover:bg-transparent border-zinc-800">
                        <Table.Head
                            class="text-zinc-500 font-normal text-center w-[25%]"
                            >key</Table.Head
                        >
                        <Table.Head
                            class="text-zinc-500 font-normal text-center w-[25%]"
                            >value</Table.Head
                        >
                        <Table.Head
                            class="text-zinc-500 font-normal text-center w-[25%]"
                            >expiry</Table.Head
                        >
                        <Table.Head
                            class="text-zinc-500 font-normal text-center w-[25%]"
                            >actions</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each entries as entry}
                        <Table.Row class="border-zinc-800 hover:bg-zinc-900/50">
                            <Table.Cell
                                class="font-medium text-zinc-100 text-center"
                                >{entry.key}</Table.Cell
                            >
                            <Table.Cell
                                class="text-zinc-300 break-all whitespace-normal text-center"
                                >{entry.value}</Table.Cell
                            >
                            <Table.Cell
                                class="text-zinc-300 break-all whitespace-normal text-center"
                                >{entry.expiry === -1
                                    ? "not set"
                                    : entry.expiry}</Table.Cell
                            >
                            <Table.Cell>
                                <div
                                    class="flex items-center gap-4 justify-center"
                                >
                                    <EditEntryDialog
                                        value={entry.value}
                                        expiry={entry.expiry}
                                        key={entry.key}
                                    />
                                    <DeleteEntryDialog key={entry.key} />
                                </div>
                            </Table.Cell>
                        </Table.Row>
                    {/each}
                </Table.Body>
            </Table.Root>
        </div>
    {/if}
</div>
