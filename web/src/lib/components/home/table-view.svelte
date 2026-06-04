<script lang="ts">
	import { getAllQuery } from "$lib/queries/queries";
	import { Loader2, AlertCircle, Inbox } from "@lucide/svelte";
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
	class="h-14 border-b border-border px-6 flex items-center justify-between bg-background/80 backdrop-blur shrink-0"
>
	<div class="flex items-center gap-3">
		<h2 class="font-serif text-lg leading-none">entries</h2>
		{#if getAll.isSuccess}
			<span class="text-xs text-muted-foreground">
				{entries.length} {entries.length === 1 ? "key" : "keys"}
			</span>
		{/if}
		<span class="relative inline-flex size-2 ml-1" title="live">
			{#if getAll.isFetching}
				<span
					class="absolute inline-flex h-full w-full animate-ping rounded-full bg-primary/40 opacity-75"
				></span>
			{/if}
			<span class="relative inline-flex size-2 rounded-full bg-primary"></span>
		</span>
	</div>
	<div class="flex items-center gap-2">
		<Seed />
		<FlushAllDialog />
		<CreateEntryDialog />
	</div>
</header>

<div class="flex-1 overflow-y-auto p-6">
	{#if getAll.isPending}
		<div class="flex items-center justify-center gap-2 text-muted-foreground py-20">
			<Loader2 class="size-4 animate-spin" />
			<span class="text-sm">loading entries…</span>
		</div>
	{:else if getAll.isError}
		<div
			class="flex flex-col items-center justify-center gap-2 text-destructive py-20"
		>
			<AlertCircle class="size-5" />
			<p class="text-sm font-medium">couldn't load data</p>
			<p class="text-xs text-muted-foreground">
				{getAll.error?.message ?? "unknown error"}
			</p>
		</div>
	{:else if entries.length === 0}
		<div
			class="flex flex-col items-center justify-center text-center gap-3 py-20 rounded-2xl border border-dashed border-border bg-card/40"
		>
			<div class="size-12 rounded-2xl bg-accent/60 grid place-items-center text-primary">
				<Inbox class="size-5" />
			</div>
			<p class="text-sm font-medium">no entries yet</p>
			<p class="text-xs text-muted-foreground max-w-[16rem]">
				add a key with the button above, or hit seed to populate sample data
			</p>
		</div>
	{:else}
		<div class="rounded-2xl border border-border bg-card overflow-hidden shadow-sm">
			<Table.Root>
				<Table.Header>
					<Table.Row class="border-border bg-muted/40 hover:bg-muted/40">
						<Table.Head class="h-10 px-4 text-[11px] font-medium uppercase tracking-wider text-muted-foreground">
							key
						</Table.Head>
						<Table.Head class="h-10 px-4 text-[11px] font-medium uppercase tracking-wider text-muted-foreground">
							value
						</Table.Head>
						<Table.Head class="h-10 px-4 text-[11px] font-medium uppercase tracking-wider text-muted-foreground w-32">
							ttl
						</Table.Head>
						<Table.Head class="h-10 px-4 text-[11px] font-medium uppercase tracking-wider text-muted-foreground text-right w-40">
							actions
						</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each entries as entry (entry.key)}
						<Table.Row class="border-border hover:bg-accent/30 transition-colors">
							<Table.Cell class="px-4 py-3 font-mono text-xs text-foreground">
								{entry.key}
							</Table.Cell>
							<Table.Cell class="px-4 py-3 font-mono text-xs text-muted-foreground break-all whitespace-normal max-w-md">
								{entry.value}
							</Table.Cell>
							<Table.Cell class="px-4 py-3 text-xs">
								{#if entry.expiry === -1}
									<span class="text-muted-foreground">∞</span>
								{:else}
									<span class="font-mono tabular-nums text-foreground">{entry.expiry}s</span>
								{/if}
							</Table.Cell>
							<Table.Cell class="px-4 py-3">
								<div class="flex items-center gap-1.5 justify-end">
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
